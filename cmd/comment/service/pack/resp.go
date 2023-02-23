package pack

import (
	"errors"

	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

/* 根据状态码来拼接基础的响应报文（包括状态码和信息）--这里完成一个RPC响应类型的完整包装*/

func CommentActionResp(err errno.ErrNo, comment *CommentServer.Comment) *CommentServer.DouyinCommentActionResponse {
	resp := new(CommentServer.DouyinCommentActionResponse) //创建一个大小等同于响应格式的空间
	resp.BaseResp = &CommentServer.BaseResp{               //指定响应状态，与errno种返回的错误一致
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.Comment = comment
	return resp //返回整个响应报文
}

func getCommentListResp(err errno.ErrNo, comments []*CommentServer.Comment) *CommentServer.DouyinCommentListResponse {
	resp := new(CommentServer.DouyinCommentListResponse)
	resp.BaseResp = &CommentServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.CommentList = comments
	return resp
}

// 报文的封装过程
func BuildcommentActionResp(err error, comment *CommentServer.Comment) *CommentServer.DouyinCommentActionResponse {
	if err == nil {
		return CommentActionResp(errno.Success, comment)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return CommentActionResp(e, nil)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return CommentActionResp(s, nil)
}

func BuildgetCommentListResp(err error, users []*CommentServer.Comment) *CommentServer.DouyinCommentListResponse {
	if err == nil {
		return getCommentListResp(errno.Success, users)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getCommentListResp(e, nil)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return getCommentListResp(s, nil)
}
