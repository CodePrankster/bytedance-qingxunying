package mysql

import "dousheng-backend/model"

// MQVideoListById 根据视频id做批量查询
func MQVideoListById(ids []string) ([]*model.Video, error) {
	var videoList []*model.Video
	if err := db.Where("id in (?)", ids).Find(&videoList).Error; err != nil {
		return videoList, err
	}
	return videoList, nil
}
