package common

// FavoriteActionRequest 点赞请求参数
type FavoriteActionRequest struct {
	UserId     int64  `json:"user_id"`     // 用户id
	VideoId    int64  `json:"video_id"`    // 视频id
	Token      string `json:"token"`       // 用户鉴权
	ActionType int32  `json:"action_type"` // 点赞类型  1-点赞，2-取消点赞
}
