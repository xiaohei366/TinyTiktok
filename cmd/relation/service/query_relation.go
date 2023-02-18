package service

import (
	"context"

	"github.com/xiaohei366/TinyTiktok/cmd/relation/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/relation/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type QueryRelationService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewQueryRelationService(ctx context.Context) *QueryRelationService {
	return &QueryRelationService{ctx: ctx}
}


func (s *QueryRelationService) QueryRelation(relation db.Follow) (bool, error) {
	//查询两者关系
	ok, err := dal.QueryFollowInfo(s.ctx, relation.UserID, relation.ToUserID)
	if err != nil {
		return false, errno.QueryFollowErr
	} 	
	return ok, nil
}
