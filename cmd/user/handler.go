package main

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/pack"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *UserServer.DouyinUserRegisterRequest) (resp *UserServer.DouyinUserRegisterResponse, err error) {

	//检验格式
	if len(req.Password) < int(1) || len(req.Username) < int(1) || len(req.Password) > int(32) || len(req.Username) > int(32) {
		resp = pack.BuilduserRegisterResp(errno.ParamErr, -1) // -1 代表 用户有问题
		return resp, nil
	}
	//完成注册动作
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err, -1) // -1 代表 用户有问题
		return resp, nil
	}
	var ua db.User
	//获得自增键后的id
	ua, err = service.NewGetUserService(ctx).GetUserByName(req.Username)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err, -1) // -1 代表 用户有问题
		return resp, nil
	}
	resp = pack.BuilduserRegisterResp(errno.Success, ua.Id)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *UserServer.DouyinUserLoginRequest) (resp *UserServer.DouyinUserLoginResponse, err error) {

	//检验格式
	if len(req.Password) < int(1) || len(req.Username) < int(1) || len(req.Password) > int(32) || len(req.Username) > int(32) {
		resp = pack.BuilduserLoginResp(errno.ParamErr, -1)
		return resp, nil
	}
	//获得用户的id
	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserLoginResp(err, -1)
		return resp, nil
	}
	//若响应成功
	resp = pack.BuilduserLoginResp(errno.Success, uid)
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *UserServer.DouyinUserRequest) (resp *UserServer.DouyinUserResponse, err error) {
	//调用服务逻辑
	u, err := service.NewGetUserService(ctx).GetUserById(req.UserId)
	if err != nil {
		resp = pack.BuilduserInfoResp(err, nil)
		return resp, nil
	}
	//正常则拼接信息返回响应
	resp = pack.BuilduserInfoResp(errno.Success, &u)
	return resp, nil
}

// MGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUserInfo(ctx context.Context, req *UserServer.DouyinMUserRequest) (resp *UserServer.DouyinMUserResponse, err error) {
	//正常则拼接信息返回响应
	//调用服务逻辑
	users, err := service.NewGetUserService(ctx).GetMUserById(req.UserId)
	if err != nil {
		resp = pack.BuildMuserInfoResp(err, nil)
		return resp, nil
	}
	resp = pack.BuildMuserInfoResp(errno.Success, users)
	return resp, nil
}

// ChangeUserFollowCount implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserFollowCount(ctx context.Context, req *UserServer.DouyinChangeUserFollowRequest) (resp *UserServer.BaseResp, err error) {
	//调用服务逻辑
	err = service.NewChangeUserFollowCountService(ctx).ChangeUserFollowCount(req.UserId, req.ToUserId, req.IsFollow)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//正常则拼接信息返回响应
	resp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
