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

// UpdateVideoById 根据id更新数据
func UpdateVideoById(id, num string) error {
	var video model.Video
	if err := db.Model(&video).Where("id = ?", id).Update("favorite_count", num).Error; err != nil {
		return err
	}
	return nil
}
