package errno

/* 一人一百个错误码的空间 */

// 基本错误码
const (
	SuccessCode int32 = iota + 0
	ServiceErrCode
	ParamErrCode
	FuncErrCode
)

// user模块错误码
const (
	UserAlreadyExistErrCode int32 = iota + 10001
	UserNotExistErrCode
	AuthorizationFailedErrCode
	StructConvertFailedErrCode
	ChangeUserFollowCountErrCode
	RelationRPCErrCode
	FindUserErrCode
)

// follow模块错误码
const (
	//follow模块错误码
	FollowActionErrCode int32 = iota + 10101
	ActionTypeErrCode
	QueryFollowErrCode
	UserRPCErrCode
	GetFollowListErrCode
	GetFollowerListErrCode
	GetFollowSetErrCode
)

// comment模块错误码
const (
	CommentActionErrCode int32 = iota + 10201
	GetCommentListErrCode
)