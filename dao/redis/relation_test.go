package redis

import (
	"dousheng-backend/setting"
	"fmt"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	// 配置初始化
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// redis初始化
	if err := Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	m.Run()
}

func TestRelationActionFollow(t *testing.T) {
	_, err := RelationActionFollow("123", "5")
	if err != nil {
		return
	}
}

func TestRelationActionUnFollow(t *testing.T) {
	_, err := RelationActionUnFollow("123", "5")
	if err != nil {
		return
	}
}

func TestGetFollowList(t *testing.T) {
	err, result := GetFollowList("123")
	if err != nil {
		return
	}
	for _, s := range result {
		fmt.Printf(s)
	}
}

func TestGetFollowerList(t *testing.T) {
	err, result := GetFollowerList("5")
	if err != nil {
		return
	}
	for _, s := range result {
		fmt.Printf(s)
	}
}

func TestGetFollowCount(t *testing.T) {
	count, err := GetFollowCount("123")
	if err != nil {
		return
	}
	fmt.Printf(strconv.FormatInt(count, 10))
}

func TestGetFollowerCount(t *testing.T) {
	count, err := GetFollowerCount("5")
	if err != nil {
		return
	}
	fmt.Printf(strconv.FormatInt(count, 10))
}
