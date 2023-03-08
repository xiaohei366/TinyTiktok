package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	rs "github.com/go-redis/redis/v8"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/rpc"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/service/pack"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/UserServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"gorm.io/gorm"
)

type MGetCommentService struct {
	ctx context.Context
}

func NewMGetCommentService(ctx context.Context) *MGetCommentService {
	return &MGetCommentService{ctx: ctx}
}

// 获得评论列表列表
func (s *MGetCommentService) MGetComment(videoID int64) ([]*CommentServer.Comment, error) {
	UserIDs := make([]int64, 0)
	texts := make([]string, 0)
	comments := make([]*db.Comment, 0)
	idds, _ := redis.VedioToComment.ZRevRangeByScore(redis.Ctx, strconv.Itoa(int(videoID)), &rs.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	times, err := redis.VedioToComment.ZRevRangeByScoreWithScores(redis.Ctx, strconv.Itoa(int(videoID)), &rs.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err == nil && len(idds) > 1 {
		for _, id := range idds {
			UserID, _ := redis.UserList.Get(redis.Ctx, id).Result()
			tem, _ := strconv.Atoi(UserID)
			UserIDs = append(UserIDs, int64(tem))
			text1, _ := redis.CommentList1.Get(redis.Ctx, id).Result()
			text2, _ := redis.CommentList2.Get(redis.Ctx, id).Result()
			texts = append(texts, text1+text2)
		}
	}
	if len(UserIDs) > 1 && len(texts) > 1 && err == nil && len(texts) == len(idds) && len(UserIDs) == len(idds) {
		for idx := range idds {
			id, _ := strconv.Atoi(idds[idx])
			u := &db.Comment{
				Model: gorm.Model{
					ID:        uint(id),
					CreatedAt: time.Unix(0, int64(times[idx].Score)),
				},
				Video_id:     videoID,
				Comment_text: texts[idx],
				User_id:      UserIDs[idx],
			}
			comments = append(comments, u)
		}
	} else {
		//先取出所有评论的ID
		comments, err = dal.MGetCommentList(s.ctx, videoID)
		if err != nil {
			return nil, errno.GetCommentListErr
		}
		for _, c := range comments {
			UserIDs = append(UserIDs, c.User_id)
			redis.AddComment(c.Video_id, c.User_id, c.Comment_text, int64(c.Model.ID), c.Model.CreatedAt)
		}
		//如果没有评论， 就直接返回即可
		if len(UserIDs) == 0 {
			return nil, nil
		}
	}
	//随后通过RPC 由这些ID获得 用户信息
	users, err := rpc.MGetUserInfo(s.ctx, &UserServer.DouyinMUserRequest{
		UserId: UserIDs,
	})
	if err != nil {
		return nil, errno.UserRPCErr
	}
	//挨个转换结构体---并同时封装好users信息
	r_comments := make([]*CommentServer.Comment, 0)

	for i, c := range comments {

		r_comments = append(r_comments, pack.CommentInfoConvert(users[i], c))
	}
	if len(comments) != len(r_comments) {
		return nil, errno.StructConvertFailedErr
	}
	return r_comments, nil
}
