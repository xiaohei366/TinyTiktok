package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"io"

	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *UserServer.DouyinUserLoginRequest) (int64, error) {
	//生成暗文密码
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	u, err := dal.GetUserByName(s.ctx, userName)
	if err != nil {
		return -1, err
	}

	if (u.Name != userName) || (u.Password != passWord) {
		return 0, errno.AuthorizationFailedErr
	}

	return int64(u.Id), nil
}
