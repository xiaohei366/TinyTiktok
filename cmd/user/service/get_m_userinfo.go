package service

import (
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 根据userId获得TableUser对象--这里不处理"是否关注(都默认false)"，放到relation模块处理
func (s *GetUserService) GetMUserById(ids []int64) ([]*db.User, error) {
	var u []*db.User
	for _, id := range ids {
		userInfo, err := s.GetUserById(id)
		if err != nil {
			return nil, errno.FindUserErr
		}
		u = append(u, &userInfo)
	}
	return u, nil
}
