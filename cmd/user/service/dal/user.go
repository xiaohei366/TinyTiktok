package dal

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"gorm.io/gorm"
)

// 将用户信息插到表中
func CreateUser(ctx context.Context, user *db.User) error {
	if err := db.DB.WithContext(ctx).Create(user).Error; err != nil {
		klog.Info("数据库插入出现问题")
		return err
	}
	return nil
}

// 根据userId获得User对象--一对一
func GetUserById(ctx context.Context, userId int64) (db.User, error) {
	res := db.User{}
	if err := db.DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		klog.Info("数据库由id查询User失败！")
		return res, err
	}
	return res, nil
}

// 根据username获得User对象--一对一
func GetUserByName(ctx context.Context, userName string) (db.User, error) {
	res := db.User{}
	if err := db.DB.WithContext(ctx).Where("name = ?", userName).Find(&res).Error; err != nil {
		klog.Info("数据库由name查询User失败！")
		return res, err
	}
	return res, nil
}

// 批量根据ID获得用户信息---一对多
func GetUserInfoListById(ctx context.Context, userIDs []int64) ([]*db.User, error) {
	res := make([]*db.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := db.DB.WithContext(ctx).Where("id IN ?", userIDs).Find(&res).Error; err != nil {
		klog.Info("数据库由id查询多个User失败！")
		return nil, err
	}
	return res, nil
}

// 关注操作--增加用户的关注数&增加被关注用户的粉丝数
func IncreaseFollowCount(ctx context.Context, userID int64, toUserID int64) error {
	// 事务操作,保持连贯,一个完整的关注操作
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新follow_count
		res := tx.Model(&db.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		if res.Error != nil {
			klog.Info("服务器增加follow_count失败")
			return res.Error
		}
		// 更新 follower_count 字段
		res = tx.Model(&db.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		if res.Error != nil {
			klog.Info("服务器增加follower_count失败")
			return res.Error
		}
		return nil
	})
	if err != nil {
		klog.Info("关注操作的事务出现问题！")
		return err
	}
	return nil
}

// 取关操作--减少用户的关注数&减少被关注用户的粉丝数
func DecreaseFollowCount(ctx context.Context, userID int64, toUserID int64) error {
	// 事务操作,保持连贯,一个完整的关注操作
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新follow_count
		res := tx.Model(&db.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		if res.Error != nil {
			klog.Info("服务器减少follow_count失败")
			return res.Error
		}
		// 更新 follower_count 字段
		res = tx.Model(&db.User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		if res.Error != nil {
			klog.Info("服务器减少follower_count失败")
			return res.Error
		}
		return nil
	})
	if err != nil {
		klog.Info("取关操作的事务出现问题！")
		return err
	}
	return nil
}
