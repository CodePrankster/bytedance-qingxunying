package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"strconv"
)

type VideoListInfo struct {
	*common.Response
	VideoList []*common.Video `json:"video_list"`
}

func NewVideoListInfo() *VideoListInfo {
	return &VideoListInfo{}
}
func FavoriteAction(request *common.FavoriteActionRequest) (int32, error) {
	uid := strconv.Itoa(int(request.UserId))
	vid := strconv.Itoa(int(request.VideoId))
	actionType := request.ActionType

	// redis操作点赞
	return redis.FavoriteAction(uid, vid, actionType)

}
func (f *VideoListInfo) FavoriteList(request *common.FavoriteListRequest) (error, *VideoListInfo) {
	uid := strconv.Itoa(int(request.UserId))
	// TODO 参数校验

	// 拿到当前用户的所有点赞视频id
	err, ids := redis.FavoriteList(uid)
	if err != nil {
		f.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return err, nil
	}
	// 根据ids查询出所有视频信息
	videoList, err := mysql.MQVideoListById(ids)
	if err != nil {
		f.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return err, nil
	}

	uids := make([]uint, len(videoList))
	for _, video := range videoList {
		uid := video.Uid
		uids = append(uids, uid)
	}

	videoInfos := make([]*common.Video, 0)

	for _, video := range videoList {
		//查询用户基本信息
		author, err := GetUserBaseInfo(video.Uid, strconv.Itoa(int(video.Uid)))

		if err != nil {
			if err != nil {
				return err, nil
			}
		}
		vid := strconv.Itoa(int(video.ID))
		num, err := redis.GetVideoFavoriteNum(vid)
		if err != nil {
			return err, nil
		}

		uid := strconv.Itoa(int(request.UserId))
		IsFavorite, _ := redis.IsFavorite(uid, vid)
		commentNum, err := mysql.GetVideoCommentNum(int64(video.ID))

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
	f.VideoList = videoInfos
	f.Response = &common.Response{
		StatusCode: common.SUCCESS,
		StatusMsg:  common.GetMsg(common.SUCCESS),
	}
	return nil, f

}
