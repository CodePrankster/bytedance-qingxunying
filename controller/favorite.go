package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/model"
)

type VideoListResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list"`
}
