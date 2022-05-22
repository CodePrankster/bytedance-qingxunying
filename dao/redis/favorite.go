package redis

import (
	"dousheng-backend/common"
	"fmt"
	"github.com/go-redis/redis"
)

// FavoriteAction 点赞具体逻辑
func FavoriteAction(uid, vid string, actionType int32) (int32, error) {

	// 首先判断当前用户是否点赞
	value := client.ZScore(GetRedisKey(KeyVideoFavoriteZSetPF+vid), uid).Val()

	pipeline := client.TxPipeline()
	if actionType == 1 && value != 1 {
		// 点赞成功，添加数据
		pipeline.ZAdd(GetRedisKey(KeyVideoFavoriteZSetPF+vid), redis.Z{
			Score:  float64(actionType),
			Member: uid,
		})
		// 记录用户点赞的视频
		pipeline.SAdd(GetRedisKey(KeyUserSetPF+uid), vid)
	}
	if actionType == 2 && value == 1 {
		// 取消点赞，把赞的类型变为2
		//client.ZRem(GetRedisKey(KeyVideoFavoriteZSetPF+vid), uid)
		pipeline.ZAdd(GetRedisKey(KeyVideoFavoriteZSetPF+vid), redis.Z{
			Score:  float64(actionType),
			Member: uid,
		})
		// 删除set里面的数据
		pipeline.SRem(GetRedisKey(KeyUserSetPF+uid), vid)
	}
	_, err := pipeline.Exec()
	if err != nil {
		return common.ERROR, err
	}

	return common.SUCCESS, nil
}

// FavoriteList 当前用户的点赞列表
func FavoriteList(uid string) (error, []string) {

	key := GetRedisKey(KeyUserSetPF + uid)
	result, err := client.SMembers(key).Result()
	if err != nil {
		return err, nil
	}
	fmt.Printf("%v\n", result)
	return nil, result

}

// 查询视频点赞总数
func GetVideoFavoriteNum(id string) (int64, error) {
	key := GetRedisKey(KeyVideoFavoriteZSetPF + id)
	cmder := client.ZCount(key, "1", "1")
	num, err := cmder.Result()
	if err != nil {
		return num, err
	}
	fmt.Printf("数量为：%d\n", num)
	return num, nil
}
