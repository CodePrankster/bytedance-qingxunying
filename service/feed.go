package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"dousheng-backend/util"
	"strconv"
	"time"
)

func FeedList(request *common.FeedRequest) (common.FeedResponse, error) {

	t, _ := time.Parse("2006-01-02 15:04:05", request.LatestTime)
	if t.IsZero() || time.Now().Sub(t) < 0 {
		// 说明传入时间为空或者传入时间大于当前时间，那么赋值为当前时间
		t = time.Now()
	}
	// 根据时间查找视频
	videos, err := mysql.GetVideoFeedListByLatestTime(t)
	if err != nil {
		msg := "根据时间查找视频操作失败"
		return common.FeedResponse{
			NextTime:   nil,
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, err
	}
	var NextTime int64
	if len(videos)-1 >= 0 {
		NextTime = videos[len(videos)-1].UpdatedAt.Unix()
	}
	// 根据视频信息查询到用户信息，并将用户信息封装到videoInfo中
	userId, _ := util.TokenVerify(request.Token)
	videoInfos, err := GetVideoListVideoIds(videos, userId)

	if err != nil {
		msg := "查找视频详细信息操作失败"
		return common.FeedResponse{
			NextTime:   nil,
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  nil,
		}, err
	}

	msg := "查找视频流操作成功"
	return common.FeedResponse{
		NextTime:   &NextTime,
		StatusCode: 0,
		StatusMsg:  &msg,
		VideoList:  videoInfos,
	}, nil

}

func GetVideoListVideoIds(videos []*model.Video, userId uint) ([]common.Video, error) {

	videoInfos := make([]common.Video, 0)

	for _, video := range videos {

		//查询用户基本信息
		author, err := GetUserBaseInfo(video.Uid, strconv.Itoa(int(video.Uid)))
		if err != nil {
			return nil, err
		}
		//查询视频的基本信息
		vid := strconv.Itoa(int(video.ID))

		num, err := redis.GetVideoFavoriteNum(vid)
		if err != nil {
			return nil, err
		}

		// 评论数量
		commentNum, err := mysql.GetVideoCommentNum(int64(video.ID))
		if err != nil {
			return nil, err
		}

		uid := strconv.Itoa(int(userId))
		IsFavorite, _ := redis.IsFavorite(uid, vid)

		videoInfos = append(videoInfos, common.Video{
			ID:            int64(video.ID),
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: num,
			CommentCount:  commentNum,
			IsFavorite:    IsFavorite,
			Title:         video.Title,
			Author:        author,
		})
	}
	return videoInfos, nil
}
