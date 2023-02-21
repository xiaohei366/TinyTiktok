package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type GetMUserService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewGetMUserService(ctx context.Context) *GetMUserService {
	return &GetMUserService{ctx: ctx}
}

// 根据userId获得TableUser对象--这里不处理"是否关注(都默认false)"，放到relation模块处理
func (s *GetMUserService) GetMUserById(ids []int64) ([]*db.User, error) {
	u, err := dal.GetUserInfoListById(s.ctx, ids)
	if err != nil {
		return nil, errno.FindUserErr
	}
	return u, nil
}
