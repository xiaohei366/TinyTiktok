package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type RegisterService struct {
	ctx context.Context
}

// 该函数可以创建一个新的注册服务--每次传入都是新的
func NewCreateUserService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// CreateUser create user info.
func (s *RegisterService) CreateUser(req *UserServer.DouyinUserRegisterRequest) error {
	//先查询有无此账号
	u, err := dal.GetUserByName(s.ctx, req.Username)
	if err != nil {
		return err
	}
	//存在则返回已存在的错误---返回的是pkg中的定义好的错误信息->会在handler调用pack集中处理好格式
	if req.Username == u.Name {
		return errno.UserAlreadyExistErr
	}
	//使用md5进行密码加密，数据库不明文存放密码
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	//完成写入账号操作
	return dal.CreateUser(s.ctx, &db.User{
		Name:     req.Username,
		Password: password,
	})
}
