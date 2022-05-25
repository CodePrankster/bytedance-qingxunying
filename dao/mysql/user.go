package mysql

import (
	"dousheng-backend/model"
	"fmt"
)

func MQueryUserById(ids []uint) (map[uint]*model.User, error) {
	var users []*model.User
	if err := db.Where("id in (?)", ids).Find(&users).Error; err != nil {
		fmt.Println("批量查询失败")
		return nil, err
	}
	// 做成索引存到内存
	userMap := make(map[uint]*model.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	return userMap, nil
}

// MQUserListById 根据用户id做批量查询
func MQUserListById(ids []string) ([]*model.User, error) {
	var userList []*model.User
	if err := db.Where("id in (?)", ids).Find(&userList).Error; err != nil {
		return userList, err
	}
	return userList, nil
}
