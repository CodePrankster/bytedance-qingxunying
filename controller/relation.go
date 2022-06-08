package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 关注功能
func RelationAvtion(c *gin.Context) {
	// 参数校验，只能支持登录的用户关注
	// 判断用户是否登录

	// 参数解析
	request := new(common.RelationActionRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	code, err := service.RelationAction(request)
	if err != nil {
		common.Error(c, code)
		return
	}
	common.OK(c, code)
	return
}

// 关注列表查询
func FollowList(c *gin.Context) {
	// 参数校验
	// 参数解析
	request := new(common.RelationFollowListRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	err, res := service.NewUserListInfo().FollowList(request)
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// 粉丝列表查询
func FollowerList(c *gin.Context) {
	// 参数校验
	// 参数解析
	request := new(common.RelationFollowerListRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	err, res := service.NewUserListInfo().FollowerList(request)
	if err != nil {
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
