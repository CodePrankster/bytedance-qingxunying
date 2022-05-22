package redis

import (
	"dousheng-backend/common"
	"github.com/go-redis/redis"
)

// FavoriteAction 点赞具体逻辑
func FavoriteAction(uid, vid string, actionType int32) (int32, error) {

	// 首先判断当前用户是否点赞
	value := client.ZScore(GetRedisKey(KeyVideoFavoriteZSetPF+vid), uid).Val()
	if actionType == 1 && value != 1 {
		// 点赞成功，添加数据
		client.ZAdd(GetRedisKey(KeyVideoFavoriteZSetPF+vid), redis.Z{
			Score:  float64(actionType),
			Member: uid,
		})
	}
	if actionType == 2 && value == 1 {
		// 取消点赞，把赞的类型变为2
		//client.ZRem(GetRedisKey(KeyVideoFavoriteZSetPF+vid), uid)
		client.ZAdd(GetRedisKey(KeyVideoFavoriteZSetPF+vid), redis.Z{
			Score:  float64(actionType),
			Member: uid,
		})
	}
	if actionType == 2 && value != 1 {

	}
	return common.SUCCESS, nil
}
