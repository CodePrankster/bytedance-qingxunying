package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/model"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list"`
}

// FavoriteAction 视频点赞功能
func FavoriteAction(c *gin.Context) {
	// 参数校验，只能支持登录的用户点赞
	// 判断用户是否登录

	// 参数解析
	request := new(common.FavoriteActionRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	code, err := service.FavoriteAction(request)
	if err != nil {
		common.Error(c, code)
		return
	}
	common.OK(c, code, nil)
	return
}
