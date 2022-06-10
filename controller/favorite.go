package controller

import (
	"dousheng-backend/common"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteAction 视频点赞功能
func FavoriteAction(c *gin.Context) {
	// 参数校验，只能支持登录的用户点赞
	// 判断用户是否登录
	uid, ok := c.Get("userId")
	if !ok {
		fmt.Println(common.GetMsg(common.CodeNeedLogin))
		zap.L().Error("没有登录")
		common.Error(c, common.CodeNeedLogin)
		return
	}
	// 参数解析
	request := new(common.FavoriteActionRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		zap.L().Error("参数解析失败", zap.Error(err))
		return
	}
	request.UserId = uid.(uint)
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
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)
	res, err := service.FavoriteList(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
