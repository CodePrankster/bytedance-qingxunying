package redis

import (
	"dousheng-backend/common"
)

func RelationActionFollow(userId, toUserId string) (int32, error) {

	// 获取redisKey
	followKey := GetRedisKey(KeyUserFollowPF + userId)
	followerKey := GetRedisKey(KeyUserFollowerPF + toUserId)

	pipeline := client.TxPipeline()

	//将被关注者加入关注者的关注列表
	pipeline.SAdd(followKey, toUserId)

	//将关注者加入被关注者的粉丝列表
	pipeline.SAdd(followerKey, userId)

	_, err := pipeline.Exec()
	if err != nil {
		return common.ERROR, err
	}

	return common.SUCCESS, nil
}

func RelationActionUnFollow(userId, toUserId string) (int32, error) {

	// 获取redisKey
	followKey := GetRedisKey(KeyUserFollowPF + userId)
	followerKey := GetRedisKey(KeyUserFollowerPF + toUserId)

	pipeline := client.TxPipeline()

	//将被关注者从关注者的关注列表中删除
	pipeline.SRem(followKey, toUserId)

	//将关注者从被关注者的粉丝列表中删除
	pipeline.SRem(followerKey, userId)

	_, err := pipeline.Exec()
	if err != nil {
		return common.ERROR, err
	}

	return common.SUCCESS, nil
}

// GetFollowList 获取用户的关注列表
func GetFollowList(userId string) (error, []string) {
	key := GetRedisKey(KeyUserFollowPF + userId)
	result, err := client.SMembers(key).Result()
	if err != nil {
		return err, nil
	}
	return nil, result
}

// GetFollowerList 获取用户的粉丝列表
func GetFollowerList(userId string) (error, []string) {
	key := GetRedisKey(KeyUserFollowerPF + userId)
	result, err := client.SMembers(key).Result()
	if err != nil {
		return err, nil
	}
	return nil, result
}

// GetFollowCount 查询关注总数
func GetFollowCount(userId string) (int64, error) {
	key := GetRedisKey(KeyUserFollowPF + userId)
	result, err := client.SCard(key).Result()
	if err != nil {
		return result, err
	}
	return result, err
}

// GetFollowerCount 查询粉丝总数
func GetFollowerCount(userId string) (int64, error) {
	key := GetRedisKey(KeyUserFollowerPF + userId)
	result, err := client.SCard(key).Result()
	if err != nil {
		return result, err
	}
	return result, err
}

// IsFollow 判断是否关注
func IsFollow(userId, toUserId string) (bool, error) {

	// 获取redisKey
	followKey := GetRedisKey(KeyUserFollowPF + userId)

	result, err := client.SIsMember(followKey, toUserId).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}
