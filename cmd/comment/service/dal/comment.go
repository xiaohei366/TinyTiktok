package dal

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/xiaohei366/TinyTiktok/cmd/comment/initialize/db"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/CommentServer"
	"gorm.io/gorm"
)

// 增加评论操作
func AddComment(ctx context.Context, comment *db.Comment) (*db.Comment, error) {

	err := db.DB.WithContext(ctx).Create(&comment).Error
	if err != nil {
		klog.Info("数据库增加出现问题")
		return comment, err
	}

	return comment, nil
}

// 删除关系操作
func DelComment(ctx context.Context, req *CommentServer.DouyinCommentActionRequest) error {
	// 需不需要检验是否为自己的id （用户只能删除自己的评论）采用事务处理机制
	err := db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		comment := new(db.Comment)
		err := tx.Where("video_id = ? And comment_id = ?", req.VideoId, req.CommentId).Find(&comment).Error
		if err != nil {
			klog.Info("不存在此评论！")
			return err
		}
		//不使用软删-防止占空间
		err = tx.Unscoped().Delete(&comment).Error
		if err != nil {
			klog.Info("删除评论失败！")
			return err
		}
		return nil
	})
	if err != nil {
		klog.Info("删除评论的事务出现问题！")
		return err
	}
	return nil
}

// 获取评论列表
func MGetCommentList(ctx context.Context, videoID int64) ([]*db.Comment, error) {
	fmt.Println("列表数据库操作")

	commentList := make([]*db.Comment, 0)

	err := db.DB.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at desc").Find(&commentList).Error
	if len(commentList) == 0 {
		return nil, nil
	}

	if err != nil {
		klog.Info("获取评论列表失败！")
		return nil, err
	}
	return commentList, nil
}
