package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/redis"
	"strconv"
)

func FavoriteAction(request *common.FavoriteActionRequest) (int32, error) {
	uid := strconv.Itoa(int(request.UserId))
	vid := strconv.Itoa(int(request.VideoId))
	actionType := request.ActionType

	return redis.FavoriteAction(uid, vid, actionType)
}
