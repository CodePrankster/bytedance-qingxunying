package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"strconv"
)

func FavoriteAction(request *common.FavoriteActionRequest) (int32, error) {
	uid := strconv.Itoa(int(request.UserId))
	vid := strconv.Itoa(int(request.VideoId))
	actionType := request.ActionType

	// redis操作点赞
	return redis.FavoriteAction(uid, vid, actionType)

}
func FavoriteList(userId uint) (common.PublishListAndFavoriteListResponse, error) {
	uid := strconv.Itoa(int(userId))
	// TODO 参数校验
	//1 拿到当前用户的所有点赞视频id
	err, ids := redis.FavoriteList(uid)
	if err != nil {
		msg := "查询用户所点赞视频的id操作失败"
		return common.PublishListAndFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, err
	}

	// 根据ids查询出所有视频信息
	videoBaseList, err := mysql.MQVideoListById(ids)
	if err != nil {
		msg := "查询用户所点赞视频的信息操作失败"
		return common.PublishListAndFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, err
	}

	//3 有视频封装数据返回
	videos := make([]common.Video, 0)

	for _, video := range videoBaseList {
		//查询用户基本信息
		author, _ := GetUserBaseInfo(video.Uid, strconv.Itoa(int(video.Uid)))
		//查询视频的点赞总数
		FavoriteCount, _ := redis.GetVideoFavoriteNum(string(video.ID))
		//查询视频的评论总数
		CommentCount, _ := mysql.GetVideoCommentNum(int64(video.ID))
		//查询视频是否点赞
		IsFavorite, _ := redis.IsFavorite(string(userId), string(video.ID))

		videos = append(videos, common.Video{
			Author:        author,
			CommentCount:  CommentCount,
			CoverURL:      video.CoverUrl,
			FavoriteCount: FavoriteCount,
			ID:            int64(video.ID),
			IsFavorite:    IsFavorite,
			PlayURL:       video.PlayUrl,
			Title:         video.Title,
		})
	}
	msg := "查询成功"
	return common.PublishListAndFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videos,
	}, nil

}
