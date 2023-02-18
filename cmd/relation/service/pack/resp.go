package pack

import (
	"errors"


	"github.com/xiaohei366/TinyTiktok/kitex_gen/RelationServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)
/* 根据状态码来拼接基础的响应报文（包括状态码和信息）--这里完成一个RPC响应类型的完整包装*/
//报文的封装
func relationActionResp(err errno.ErrNo) *RelationServer.DouyinRelationActionResponse {
	resp := new(RelationServer.DouyinRelationActionResponse)
	resp.BaseResp = &RelationServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg: err.ErrMsg,
	}
	return resp
}

func relationQueryResp(err errno.ErrNo, isFollow bool) *RelationServer.DouyinQueryRelationResponse {
	resp := new(RelationServer.DouyinQueryRelationResponse)
	resp.BaseResp = &RelationServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg: err.ErrMsg,
	}
	resp.IsFollow = isFollow
	return resp
}

func getFollowListResp(err errno.ErrNo, users []*RelationServer.User) *RelationServer.DouyinRelationFollowListResponse {
	resp := new(RelationServer.DouyinRelationFollowListResponse)
	resp.BaseResp = &RelationServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg: err.ErrMsg,
	}
	resp.UserList = users
	return resp
}

func getFollowerListResp(err errno.ErrNo, users []*RelationServer.User) *RelationServer.DouyinRelationFollowerListResponse {
	resp := new(RelationServer.DouyinRelationFollowerListResponse)
	resp.BaseResp = &RelationServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg: err.ErrMsg,
	}
	resp.UserList = users
	return resp
}

// 报文的封装过程
func BuildrelationActionResp(err error) *RelationServer.DouyinRelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationActionResp(s)
}

func BuildrelationQueryResp(err error, isFollow bool) *RelationServer.DouyinQueryRelationResponse {
	if err == nil {
		return relationQueryResp(errno.Success, isFollow)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationQueryResp(e, isFollow)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationQueryResp(s, isFollow)
}

func BuildgetFollowListResp(err error, users []*RelationServer.User) *RelationServer.DouyinRelationFollowListResponse {
	if err == nil {
		return getFollowListResp(errno.Success, users)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getFollowListResp(e, nil)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return getFollowListResp(s, nil)
}

func BuildgetFollowerListResp(err error, users []*RelationServer.User) *RelationServer.DouyinRelationFollowerListResponse {	
	if err == nil {
		return getFollowerListResp(errno.Success, users)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getFollowerListResp(e, nil)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return getFollowerListResp(s, nil)
}