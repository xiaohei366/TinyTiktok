package pack

import (
	"errors"

	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

/* 根据状态码来拼接基础的响应报文（包括状态码和信息）---可复用！*/
//报文的封装
func baseResp(err errno.ErrNo) *UserServer.BaseResp {
	return &UserServer.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

func userLoginResp(err errno.ErrNo, userID int64) *UserServer.DouyinUserLoginResponse {
	resp := new(UserServer.DouyinUserLoginResponse)
	resp.BaseResp = &UserServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.UserId = userID
	return resp
}

func userRegisterResp(err errno.ErrNo, userID int64) *UserServer.DouyinUserRegisterResponse {
	resp := new(UserServer.DouyinUserRegisterResponse)
	resp.BaseResp = &UserServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.UserId = userID
	return resp
}

func userInfoResp(err errno.ErrNo, user *UserServer.User) *UserServer.DouyinUserResponse {
	resp := new(UserServer.DouyinUserResponse)
	resp.BaseResp = &UserServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.User = user
	return resp
}

func MuserInfoResp(err errno.ErrNo, user []*UserServer.User) *UserServer.DouyinMUserResponse {
	resp := new(UserServer.DouyinMUserResponse)
	resp.BaseResp = &UserServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.User = user
	return resp
}

// 报文的封装过程
func BuildBaseResp(err error) *UserServer.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func BuilduserLoginResp(err error, userID int64) *UserServer.DouyinUserLoginResponse {
	if err == nil {
		return userLoginResp(errno.Success, userID)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userLoginResp(e, userID)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userLoginResp(s, userID)
}

func BuilduserRegisterResp(err error, userID int64) *UserServer.DouyinUserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success, userID)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e, userID)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userRegisterResp(s, userID)
}

func BuilduserInfoResp(err error, u *db.User) *UserServer.DouyinUserResponse {
	//先转换结构体
	user := UserInfoConvert(u)
	//随后再进行响应报文的封装
	if err == nil {
		return userInfoResp(errno.Success, user)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userInfoResp(e, user)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userInfoResp(s, user)
}

func BuildMuserInfoResp(err error, db_users []*db.User) *UserServer.DouyinMUserResponse {
	if len(db_users) == 0 {
		return MuserInfoResp(errno.Success, nil)
	}
	//先挨个转换结构体
	users := make([]*UserServer.User, 0)
	for _, u := range db_users {
		users = append(users, UserInfoConvert(u)) //由于此处不查询关系，因此一律false处理
	}
	//转换过程出现问题
	if len(users) != len(db_users) {
		return MuserInfoResp(errno.StructConvertFailedErr, nil)
	}
	//随后再进行响应报文的封装
	if err == nil {
		return MuserInfoResp(errno.Success, users)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return MuserInfoResp(e, users)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return MuserInfoResp(s, users)
}
