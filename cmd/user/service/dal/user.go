package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
)


// 将用户信息插到表中
func CreateUser(ctx context.Context, user *db.User) error {
	if err := db.DB.WithContext(ctx).Create(user).Error; err != nil {
		klog.Info("数据库插入出现问题")
		return err
	}
	return nil
}

//根据userId获得User对象--一对一
func GetUserById(ctx context.Context, userId int64) (db.User, error) {
	res := db.User{}
	if err := db.DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

//根据username获得User对象--一对一
func GetUserByName(ctx context.Context, userName string) (db.User, error) {
	res := db.User{}
	if err := db.DB.WithContext(ctx).Where("name = ?", userName).Find(&res).Error; err != nil {
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
		return nil, err
	}
	return res, nil
}