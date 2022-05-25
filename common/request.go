package common

// FavoriteActionRequest 点赞请求参数
type FavoriteActionRequest struct {
	UserId     int64  `json:"user_id"`     // 用户id
	VideoId    int64  `json:"video_id"`    // 视频id
	Token      string `json:"token"`       // 用户鉴权
	ActionType int32  `json:"action_type"` // 点赞类型  1-点赞，2-取消点赞
}

// FavoriteListRequest 点赞请求参数
type FavoriteListRequest struct {
	UserId int64 `json:"user_id" form:"user_id"` // 用户id

	Token string `json:"token" form:"token"` // 用户鉴权

}

// RelationActionRequest 关系操作参数
type RelationActionRequest struct {
	UserId     int64  `json:"user_id"`     //用户id
	Token      string `json:"token"`       // 用户鉴权
	ToUserId   int64  `json:"to_user_id"`  // 对方用户id
	ActionType int32  `json:"action_type"` // 1-关注 2-取消关注
}

// RelationFollowListRequest 关系操作参数
type RelationFollowListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"` // 用户id
	Token  string `json:"token" form:"token"`     // 用户鉴权
}

// RelationFollowerListRequest 关系操作参数
type RelationFollowerListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"` // 用户id
	Token  string `json:"token" form:"token"`     // 用户鉴权
}
