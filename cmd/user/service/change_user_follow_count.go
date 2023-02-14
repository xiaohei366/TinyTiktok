package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type ChangeUserFollowCountService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewChangeUserFollowCountService(ctx context.Context) *ChangeUserFollowCountService {
	return &ChangeUserFollowCountService{ctx: ctx}
}


// 根据userId获得TableUser对象--这里不处理"是否关注(都默认false)"，放到relation模块处理
func (s *ChangeUserFollowCountService) ChangeUserFollowCount(userID int64, toUserID int64, isFollow bool) error {
	var err error
	if isFollow {
		err = dal.IncreaseFollowCount(s.ctx, userID, toUserID)
	} else {
		err = dal.DecreaseFollowCount(s.ctx, userID, toUserID)
	}
	if err != nil {
		return errno.ChangeUserFollowCountErr
	}
	return nil
}