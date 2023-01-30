package pack

import (
	"errors"

	"github.com/xiaohei366/TinyTiktok/cmd/user/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)
/* 根据状态码来拼接基础的响应报文（包括状态码和信息）---可复用！*/
// 根据pkg中定义好的错误码变量，自动生成错误的响应code和响应message
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

//将pkg中定义好的错误码都转化到相应对象的状态码和信息
func baseResp(err errno.ErrNo) *UserServer.BaseResp {
	return &UserServer.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
