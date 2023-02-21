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

	// user模块错误码
	UserAlreadyExistErr    = NewErrNo(int32(10003), "User already exists")
	UserNotExistErr        = NewErrNo(int32(10004), "User not exists")
	AuthorizationFailedErr = NewErrNo(int32(10005), "Authorization failed")

	// favorite模块错误码
	FavoriteNotExistErr   = NewErrNo(int32(10006), "Favorite not exist")
	FavouriteActionErr    = NewErrNo(int32(10007), "FavoriteAction failed")
	FavoriteActionTypeErr = NewErrNo(int32(10008), "FavoriteActionType is wrong")
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
	//s := ServiceErr
	//s.ErrMsg = err.Error()
	return Err
}
