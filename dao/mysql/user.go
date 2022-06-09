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
func CheckUserName(username string) int8 {
	var registInfo model.RegistInfo
	db.First(&registInfo, "user_name = ?", username)
	if registInfo.UserName != "" {
		return 1
	}
	return 0
}
func SelectUser(username string, password string) *model.RegistInfo {
	var registInfo *model.RegistInfo
	db.First(&registInfo, "user_name = ? and password = ?", username, password)
	return registInfo
}

func SelectUserName(userId int64) (string, error) {
	var user *model.User
	db.Select("name").Where("id = ?", userId).First(&user)
	return user.Name, nil
}
func InsertUserInfo(info model.User) error {
	if err := db.Create(&info).Error; err != nil {
		return err
	}
	return nil
}
