package mysql

import (
	"dousheng-backend/common"
	"dousheng-backend/model"
)

func InsertComment(comment *model.Comment) (int32, error) {
	if err := db.Create(comment).Error; err != nil {
		return common.ERROR, nil
	}
	return common.SUCCESS, nil
}

func DeleteComment(vid int64) (int32, error) {
	p := new(model.Comment)
	if err := db.Where("vid = ?", vid).Delete(p).Error; err != nil {
		return common.ERROR, err
	}
	return common.SUCCESS, nil
}

func GetCommentListByVid(vid int64) (int32, []*model.Comment, error) {
	var commentList []*model.Comment
	if err := db.Where("vid = ?", vid).Find(&commentList).Error; err != nil {
		return common.ERROR, nil, err
	}
	return common.SUCCESS, commentList, nil
}
