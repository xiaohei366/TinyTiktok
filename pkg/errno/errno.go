package errno

import (
	"errors"
	"fmt"
)

// 错误码定义---统一在此定义-防止错误码冲突
var (
	Success    = NewErrNo(int32(0), "Success")
	ServiceErr = NewErrNo(int32(10001), "Service is unable to start successfully")
	ParamErr   = NewErrNo(int32(10002), "Wrong Parameter has been given")
	FuncErr    = NewErrNo(int32(10003), "Error!")
	// user模块错误码
	UserAlreadyExistErr    = NewErrNo(int32(10004), "User already exists")
	UserNotExistErr        = NewErrNo(int32(10005), "User not exists")
	AuthorizationFailedErr = NewErrNo(int32(10006), "Authorization failed")
	StructConvertFailedErr = NewErrNo(int32(10007), "Struct Convert failed")
	ChangeUserFollowCountErr = NewErrNo(int32(10008), "Failed to modify the follow count")
	RelationRPCErr = NewErrNo(int32(10009), "Failed to use relation RPC")
	FindUserErr = NewErrNo(int32(10010), "Failed to use relation RPC")
	//follow模块错误码
	FollowActionErr = NewErrNo(int32(10110), "Follow action failed")
	ActionTypeErr   = NewErrNo(int32(10111), "Wrong action-type has been given")
	QueryFollowErr  = NewErrNo(int32(10112), "Query relation failed")
	UserRPCErr  = NewErrNo(int32(10113), "Failed to use user RPC")
	GetFollowListErr = NewErrNo(int32(10114), "Failed to get follow list")
	GetFollowerListErr = NewErrNo(int32(10115), "Failed to get follower list")
	GetFollowSetErr = NewErrNo(int32(10116), "Failed to get follow set")
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// 创建错误码信息--最基础的两个点
func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// 附加修改errno的信息
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := FuncErr
	s.ErrMsg = err.Error()
	return Err
}
