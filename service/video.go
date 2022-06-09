package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"dousheng-backend/util"
	"fmt"
	"strconv"
	"time"
)

type VideoFeedListInfo struct {
	*common.Response
	NextTime  int64           `json:"next_time"`
	VideoList []*common.Video `json:"video_list"`
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
	userId, _ := util.TokenVerify(request.Token)

	videoInfos, err := GetUsersByVideoIds(videos, userId)

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

func GetUsersByVideoIds(videos []*model.Video, userId uint) ([]*common.Video, error) {
	uids := make([]uint, len(videos))
	for _, video := range videos {
		uid := video.Uid
		uids = append(uids, uid)
	}

	videoInfos := make([]*common.Video, 0)

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

		videoInfos = append(videoInfos, &common.Video{
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
