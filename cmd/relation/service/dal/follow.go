package dal

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"gorm.io/gorm"
)

// 增加关系操作
func AddFollow(ctx context.Context, userId int64, toUserId int64) error {
	follow := db.Follow{
		UserID:   userId,
		ToUserID: toUserId,
	}
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先判断是否存在关系
		ok := new(db.Follow)
		err := tx.Where("user_id = ? And to_user_id = ?", userId, toUserId).Find(&ok).Error
		if err != nil || ok.ID != 0 {
			klog.Info("数据库查询出现问题Or用户已存在")
			return err
		}
		//如果不存在，则可以插入
		if err := db.DB.WithContext(ctx).Select("user_id", "to_user_id").Create(&follow).Error; err != nil {
			klog.Info("数据库插入出现问题")
			return err
		}
		return nil
	})
	if err != nil {
		klog.Info("添加关系的事务出现问题！")
		return err
	}
	return nil
}

// 删除关系操作
func DelFollow(ctx context.Context, userId int64, toUserId int64) error {
	// 因为不止一步CRUD，因此采用事务处理机制
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先判断是否存在关系
		follow := new(db.Follow)
		err := tx.Where("user_id = ? And to_user_id = ?", userId, toUserId).Find(&follow).Error
		if err != nil {
			klog.Info("不存在关系！")
			return err
		}
		//不使用软删-防止占空间
		err = tx.Unscoped().Delete(&follow).Error
		if err != nil {
			klog.Info("删除关系失败！")
			return err
		}
		return nil
	})
	if err != nil {
		klog.Info("删除关系的事务出现问题！")
		return err
	}
	return nil
}

// 查看两人是否有关系
func QueryFollowInfo(ctx context.Context, userID int64, toUserID int64) (bool, error) {
	if userID == toUserID {
		return true, nil
	}
	follow := make([]*db.Follow, 0)
	err := db.DB.WithContext(ctx).Find(&follow, "user_id = ? and to_user_id = ?", userID, toUserID).Error
	if err != nil {
		klog.Info("查询关系失败！")
		return false, err
	}
	return len(follow) != 0, nil
}

// 获取关注列表
func MGetFollowList(ctx context.Context, userID int64) ([]*db.Follow, error) {
	followList := make([]*db.Follow, 0)
	err := db.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&followList).Error
	if err != nil {
		klog.Info("获取关注列表失败！")
		return nil, err
	}
	return followList, nil
}

// 获取粉丝列表
func MGetFollowerList(ctx context.Context, userID int64) ([]*db.Follow, error) {
	followList := make([]*db.Follow, 0)
	err := db.DB.WithContext(ctx).Where("to_user_id = ?", userID).Find(&followList).Error
	if err != nil {
		klog.Info("获取粉丝列表失败！")
		return nil, err
	}
	return followList, nil
}

// 获得关注的集合---用于Redis的改进or快速知道与粉丝是否被关注
func GetFollowSet(ctx context.Context, userID int64) (map[int64]struct{}, error) {
	//将该用户关注的人放入集合中，以便可以快速取出对比
	followSet := make(map[int64]struct{})
	followList, err := MGetFollowList(ctx, userID)
	if err != nil {
		klog.Info("获取关注列表失败！")
		return nil, err
	}
	for _, v := range followList {
		followSet[v.ToUserID] = struct{}{}
	}
	return followSet, nil
}
