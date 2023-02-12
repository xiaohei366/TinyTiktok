package service

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type GetUserService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// 根据username获得TableUser对象
func (s *GetUserService) GetUserByName(name string) (dal.User, error) {
	u, err := dal.GetUserByName(s.ctx, name)
	if err != nil {
		return u, errno.UserNotExistErr
	}
	return u, nil
}

// 根据userId获得TableUser对象
func (s *GetUserService) GetUserById(id int64) (dal.User, error) {
	u, err := dal.GetUserById(s.ctx, id)
	if err != nil {
		return u, errno.UserNotExistErr
	}
	return u, nil
}