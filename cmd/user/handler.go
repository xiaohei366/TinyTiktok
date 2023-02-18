package main

import (
	"context"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"

	"github.com/xiaohei366/TinyTiktok/cmd/user/service"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *UserServer.DouyinUserRegisterRequest) (resp *UserServer.DouyinUserRegisterResponse, err error) {
	resp = new(UserServer.DouyinUserRegisterResponse)
	//检验格式
	if len(req.Password) < int(1) || len(req.Username) < int(1) || len(req.Password) > int(32) || len(req.Username) > int(32) {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.UserId = -1 // -1 代表 用户有问题
		return resp, nil
	}
	//完成注册动作
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = -1 // -1 代表 用户有问题
		return resp, nil
	}
	var ua dal.User
	//获得自增键后的id
	ua, err = service.NewGetUserService(ctx).GetUserByName(req.Username)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = -1 // -1 代表 用户有问题
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = ua.Id
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *UserServer.DouyinUserLoginRequest) (resp *UserServer.DouyinUserLoginResponse, err error) {
	resp = new(UserServer.DouyinUserLoginResponse)
	//检验格式
	if len(req.Password) < int(1) || len(req.Username) < int(1) || len(req.Password) > int(32) || len(req.Username) > int(32) {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		resp.UserId = -1 // -1 代表 用户有问题
		return resp, nil
	}
	//获得用户的id
	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UserId = -1 // -1 代表 用户有问题
		return resp, nil
	}
	//若响应成功
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uid
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *UserServer.DouyinUserRequest) (resp *UserServer.DouyinUserResponse, err error) {
	resp = new(UserServer.DouyinUserResponse)
	//
	u, err := service.NewGetUserService(ctx).GetUserById(req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//正常则拼接信息返回响应
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = pack.UserInfo(&u)
	return resp, nil
}
