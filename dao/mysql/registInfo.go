package mysql

import "dousheng-backend/model"

// InsertVideo
func InsertRegistInfo(info model.RegistInfo) (uint, error) {
	if err := db.Create(&info).Error; err != nil {
		return 0, err
	}
	return info.Model.ID, err
}
