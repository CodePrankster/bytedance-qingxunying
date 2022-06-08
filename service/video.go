package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"fmt"
	"strconv"
	"time"
)

type VideoFeedListInfo struct {
	*common.Response
	NextTime  int64        `json:"next_time"`
	VideoList []*VideoInfo `json:"video_list"`
}

func NewVideoFeedListInfo() *VideoFeedListInfo {
	return &VideoFeedListInfo{}
}
func (v *VideoFeedListInfo) FeedList(request *common.FeedRequest) (*VideoFeedListInfo, error) {

	t, _ := time.Parse("2006-01-02 15:04:05", request.LatestTime)
	if t.IsZero() || time.Now().Sub(t) < 0 {
		// 说明传入时间为空或者传入时间大于当前时间，那么赋值为当前时间
		t = time.Now()
	}
	// 根据时间查找视频
	videos, err := mysql.GetVideoFeedListByLatestTime(t)
	if err != nil {
		v.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return nil, err
	}

	fmt.Println(len(videos) - 1)
	if len(videos)-1 >= 0 {
		v.NextTime = videos[len(videos)-1].UpdatedAt.Unix()
	}
	// 根据视频信息查询到用户信息，并将用户信息封装到videoInfo中
	videoInfos, err := GetUsersByVideoIds(videos)
	if err != nil {
		v.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return nil, err
	}
	v.VideoList = videoInfos
	v.Response = &common.Response{
		StatusCode: common.SUCCESS,
		StatusMsg:  common.GetMsg(common.SUCCESS),
	}
	return v, err

}

func GetUsersByVideoIds(videos []*model.Video) ([]*VideoInfo, error) {
	uids := make([]uint, len(videos))
	for _, video := range videos {
		uid := video.Uid
		uids = append(uids, uid)
	}

	// mysql批量查询用户
	userMap, err := mysql.MQueryUserById(uids)
	if err != nil {
		return nil, err
	}
	videoInfos := make([]*VideoInfo, 0)

	for _, video := range videos {
		user := userMap[video.Uid]
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
		IsFavorite, _ := redis.IsFavorite(string(user.ID), string(vid))

		videoInfos = append(videoInfos, &VideoInfo{
			ID:            video.ID,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: num,
			CommentCount:  commentNum, // TODO 查询评论数量
			IsFavorite:    IsFavorite,
			Title:         video.Title,
			User:          user,
		})
	}
	return videoInfos, nil
}
