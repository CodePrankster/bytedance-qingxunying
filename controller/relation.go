package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 关注功能
func RelationAvtion(c *gin.Context) {
	// 参数校验，只能支持登录的用户关注
	// 判断用户是否登录
	uid, ok := c.Get("userId")
	if !ok {
		fmt.Println(common.GetMsg(common.CodeNeedLogin))
		common.Error(c, common.CodeNeedLogin)
		return
	}
	// 参数解析
	request := new(common.RelationActionRequest)

	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	request.UserId = uid.(uint)
	userId, _ := c.Get("userId")

	code, err := service.RelationAction(request, userId.(uint))
	if err != nil {
		common.Error(c, code)
		return
	}
	common.OK(c, code)
	return
}

// 关注列表查询
func FollowList(c *gin.Context) {
	// 参数解析
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	err, res := service.FollowList(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// 粉丝列表查询
func FollowerList(c *gin.Context) {
	// 参数解析
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	err, res := service.FollowerList(uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
