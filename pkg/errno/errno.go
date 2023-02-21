package errno

import (
	"errors"
	"fmt"
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
	//// Errno格式直接返回即可
	if errors.As(err, &Err) {
		return Err
	}
	s := FuncErr
	s.ErrMsg = err.Error()
	return Err
}

// 错误码定义---统一在此定义-防止错误码冲突
var (
	Success    = NewErrNo(SuccessCode, "Success")
	ServiceErr = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr   = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	FuncErr    = NewErrNo(FuncErrCode, "Error!")
	// user模块错误码
	UserAlreadyExistErr      = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UserNotExistErr          = NewErrNo(UserNotExistErrCode, "User not exists")
	AuthorizationFailedErr   = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	StructConvertFailedErr   = NewErrNo(StructConvertFailedErrCode, "Struct Convert failed")
	ChangeUserFollowCountErr = NewErrNo(ChangeUserFollowCountErrCode, "Failed to modify the follow count")
	RelationRPCErr           = NewErrNo(RelationRPCErrCode, "Failed to use relation RPC")
	FindUserErr              = NewErrNo(FindUserErrCode, "Failed to use relation RPC")
	//follow模块错误码
	FollowActionErr    = NewErrNo(FollowActionErrCode, "Follow action failed")
	ActionTypeErr      = NewErrNo(ActionTypeErrCode, "Wrong action-type has been given")
	QueryFollowErr     = NewErrNo(QueryFollowErrCode, "Query relation failed")
	UserRPCErr         = NewErrNo(UserRPCErrCode, "Failed to use user RPC")
	GetFollowListErr   = NewErrNo(GetFollowListErrCode, "Failed to get follow list")
	GetFollowerListErr = NewErrNo(GetFollowerListErrCode, "Failed to get follower list")
	GetFollowSetErr    = NewErrNo(GetFollowSetErrCode, "Failed to get follow set")
	//video模块错误码
	PublishActionErr    = NewErrNo(PublishActionErrCode, "Publish Action failed")
	PublishListErr      = NewErrNo(PublishListErrCode, "Publish List failed")
	FeedErr             = NewErrNo(FeedErrCode, "Feed videos failed")
	VideoRpcUserErr     = NewErrNo(VideoRpcUserErrCode, "Video rpc User failed")
	VideoRpcRelationErr = NewErrNo(VideoRpcRelationErrCode, "Video rpc relation failed")
	VideoListNotFound   = NewErrNo(VideoListNotFoundErrCode, "Video List is empty")
)
