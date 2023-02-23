package pack

import (
	"errors"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/FavoriteServer"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
)

// 报文的封装
func favoriteActionResp(err errno.ErrNo) *FavoriteServer.DouyinFavoriteActionResponse {
	resp := new(FavoriteServer.DouyinFavoriteActionResponse)
	resp.BaseResp = &FavoriteServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	return resp
}

func favoriteVideoQueryResp(err errno.ErrNo, FavoriteCount int64, isFavorite bool) *FavoriteServer.DouyinVideoFavoriteResponse {
	resp := new(FavoriteServer.DouyinVideoFavoriteResponse)
	resp.BaseResp = &FavoriteServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.FavoriteCount = FavoriteCount
	resp.IsFavorite = isFavorite
	return resp
}

func favoriteUserQueryResp(err errno.ErrNo, TotalFavorited int64, FavoriteCount int64) *FavoriteServer.DouyinUserFavoriteResponse {
	resp := new(FavoriteServer.DouyinUserFavoriteResponse)
	resp.BaseResp = &FavoriteServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.TotalBeFavorite = TotalFavorited
	resp.FavoriteCount = FavoriteCount
	return resp
}

func getFavoriteListResp(err errno.ErrNo, videos []*FavoriteServer.Video) *FavoriteServer.DouyinFavoriteListResponse {
	resp := new(FavoriteServer.DouyinFavoriteListResponse)
	resp.BaseResp = &FavoriteServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.VideoList = videos

	return resp
}

// 报文的封装过程
func BuildfavoriteActionResp(err error) *FavoriteServer.DouyinFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteActionResp(s)
}

func BuildfavoriteVideoQueryResp(err error, FavoriteCount int64, isFavorite bool) *FavoriteServer.DouyinVideoFavoriteResponse {
	if err == nil {
		return favoriteVideoQueryResp(errno.Success, FavoriteCount, isFavorite)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteVideoQueryResp(e, FavoriteCount, isFavorite)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteVideoQueryResp(s, FavoriteCount, isFavorite)
}

func BuildfavoriteUserQueryResp(err error, TotalFavorited int64, FavoriteCount int64) *FavoriteServer.DouyinUserFavoriteResponse {
	if err == nil {
		return favoriteUserQueryResp(errno.Success, TotalFavorited, FavoriteCount)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteUserQueryResp(e, TotalFavorited, FavoriteCount)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return favoriteUserQueryResp(s, TotalFavorited, FavoriteCount)

}
func queryUserLikeVideoResp(err errno.ErrNo, isFav bool) *FavoriteServer.DouyinQueryFavoriteResponse {
	resp := new(FavoriteServer.DouyinQueryFavoriteResponse)
	resp.BaseResp = &FavoriteServer.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
	resp.Favorite = isFav
	return resp
}
func BuildQueryUserLikeVideoResp(err error, isFav bool) *FavoriteServer.DouyinQueryFavoriteResponse {
	if err == nil {
		return queryUserLikeVideoResp(errno.Success, isFav)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return queryUserLikeVideoResp(e, isFav)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return queryUserLikeVideoResp(s, isFav)
}

func BuildgetFavoriteListResp(err error, videos []*FavoriteServer.Video) *FavoriteServer.DouyinFavoriteListResponse {
	if err == nil {
		return getFavoriteListResp(errno.Success, videos)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return getFavoriteListResp(e, nil)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return getFavoriteListResp(s, nil)
}

func ConvertVideos(videos []*VideoServer.Video) []*FavoriteServer.Video {
	videosList := make([]*FavoriteServer.Video, 0)
	for _, v := range videos {
		videosList = append(videosList, &FavoriteServer.Video{
			Id: v.Id,
			Author: &FavoriteServer.User{
				Id:            v.Author.Id,
				Name:          v.Author.Name,
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				IsFollow:      v.Author.IsFollow,
				//TotalBeFavorite: ,//没记录这两个数的
				//FavoriteCount: ,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return videosList
}
