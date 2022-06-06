package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"strconv"
)

type VideoListInfo struct {
	*common.Response
	VideoList []*VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	ID            uint        `json:"id"`
	User          *model.User `json:"author"`                             // 用户id
	PlayUrl       string      `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl      string      `json:"cover_url,omitempty"`                // 视频封面地址
	Title         string      `json:"title,omitempty"`                    // 视频标题
	FavoriteCount int64       `json:"favorite_count,omitempty"`           // 视频的点赞总数
	CommentCount  int64       `json:"comment_count,omitempty"`            // 视频的评论总数
	IsFavorite    bool        `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
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

	// mysql批量查询用户
	userMap, err := mysql.MQueryUserById(uids)
	if err != nil {
		f.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return err, nil
	}

	videoInfos := make([]*VideoInfo, 0)
	for _, video := range videoList {
		user := userMap[video.Uid]
		vid := strconv.Itoa(int(video.ID))
		num, err := redis.GetVideoFavoriteNum(vid)
		if err != nil {
			return err, nil
		}
		videoInfos = append(videoInfos, &VideoInfo{
			ID:            video.ID,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: num,
			CommentCount:  1, // TODO 查询评论数量
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
			User:          user,
		})
	}
	f.VideoList = videoInfos
	f.Response = &common.Response{
		StatusCode: common.SUCCESS,
		StatusMsg:  common.GetMsg(common.SUCCESS),
	}
	return nil, f

}
