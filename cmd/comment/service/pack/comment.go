package pack

import (
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
)

/* 拼接用户信息--将UserServer.User插入CommentServer.Comment*/
func CommentInfoConvert(u *UserServer.User, v *db.Comment) *CommentServer.Comment {
	if u == nil {
		return nil
	}
	return &CommentServer.Comment{
		Id: v.User_id,
		User: &CommentServer.User{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false,
		},
		Content:    v.Comment_text,
		CreateDate: v.Model.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
