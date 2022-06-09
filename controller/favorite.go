package controller

import (
	"dousheng-backend/common"
	"net/http"

	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteAction 视频点赞功能
func FavoriteAction(c *gin.Context) {
	// 参数校验，只能支持登录的用户点赞
	// 判断用户是否登录

	// 参数解析
	request := new(common.FavoriteActionRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	code, err := service.FavoriteAction(request)
	if err != nil {
		common.Error(c, code)
		return
	}
	common.OK(c, code)
	return
}

// FavoriteList 视频列表查询
func FavoriteList(c *gin.Context) {
	// 参数校验
	// 参数解析
	request := new(common.FavoriteListRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	err, res := service.NewVideoListInfo().FavoriteList(request)
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
