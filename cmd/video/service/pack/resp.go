package pack

import (
	"errors"
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/VideoServer"

	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

/* 根据状态码来拼接基础的响应报文（包括状态码和信息）---可复用！*/
// 根据pkg中定义好的错误码变量，自动生成错误的响应code和响应message
// BuildBaseResp build the videoServer response store the error msg.
func BuildBaseResp(err error) *VideoServer.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	return baseResp(e)

}

// 将pkg中定义好的错误码都转化到相应对象的状态码和信息
// baseResp convert error msg into the videoServer base response msg.
func baseResp(err errno.ErrNo) *VideoServer.BaseResp {
	return &VideoServer.BaseResp{
		StatusCode: err.ErrCode, StatusMsg: err.ErrMsg,
	}
}
