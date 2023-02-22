package service

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/db"
	"github.com/xiaohei366/TinyTiktok/cmd/user/initialize/redis"
	"github.com/xiaohei366/TinyTiktok/cmd/user/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

type GetUserService struct {
	ctx context.Context
}

// GetUserService new MGetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// 根据username获得TableUser对象
func (s *GetUserService) GetUserByName(name string) (db.User, error) {
	u, err := dal.GetUserByName(s.ctx, name)
	if err != nil {
		return u, errno.UserNotExistErr
	}
	return u, nil
}

// 根据userId获得TableUser对象
func (s *GetUserService) GetUserById(id int64) (db.User, error) {
	var u db.User
	var err error
	var FollowNum, FollowerNum string
	//先在Redis上查询---多机防止热key
	//将热key分散到不同的服务器中
	rand.Seed(time.Now().UnixNano())
	Randid := rand.Intn(2) // 有几个机器是几
	switch Randid {
	case 0:
		FollowNum, _ = redis.Count1.HGet(redis.Ctx, strconv.Itoa(int(id)), redis.FollowField).Result()
		FollowerNum, _ = redis.Count1.HGet(redis.Ctx, strconv.Itoa(int(id)), redis.FollowField).Result()
	case 1:
		FollowNum, _ = redis.Count2.HGet(redis.Ctx, strconv.Itoa(int(id)), redis.FollowField).Result()
		FollowerNum, _ = redis.Count2.HGet(redis.Ctx, strconv.Itoa(int(id)), redis.FollowField).Result()
	}
	//再查询Name
	Name, _ := redis.Name.Get(redis.Ctx, strconv.Itoa(int(id))).Result()
	if Name != "" && len(FollowNum) != 0 && len(FollowerNum) != 0 && FollowNum != "0" && FollowerNum != "0" {
		FollowInt, _ := strconv.ParseInt(FollowNum, 10, 64)
		FollowerInt, _ := strconv.ParseInt(FollowerNum, 10, 64)
		u = db.User{
			Id:            id,
			Name:          Name,
			Password:      "",
			FollowCount:   FollowInt,
			FollowerCount: FollowerInt,
		}
	} else {
		u, err = dal.GetUserById(s.ctx, id)
		if err != nil {
			return u, errno.UserNotExistErr
		}
		redis.AddName(u.Id, u.Name)
		redis.UpdateCount(u.Id, u.FollowCount, u.FollowerCount)
	}
	return u, nil
}
