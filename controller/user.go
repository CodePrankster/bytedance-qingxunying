package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// 参数解析
	request := new(common.RegAndLogRequest)
	err := c.ShouldBindQuery(request)
	if err != nil {
		fmt.Println("参数解析失败")
		return
	}
	response := service.Register(request)
	c.JSON(http.StatusOK, response)
	return
}

func Login(c *gin.Context) {
	// 参数解析
	request := new(common.RegAndLogRequest)
	err := c.ShouldBindQuery(request)
	if err != nil {
		fmt.Println("参数解析失败")
		return
	}
	response := service.Login(request)
	c.JSON(http.StatusOK, response)
	return
}

//UserInfo 获取用户的信息
func UserInfo(c *gin.Context) {
	// 参数解析

	request := new(common.UserInfoRequese)
	err := c.ShouldBindQuery(request)
	if err != nil {
		fmt.Println("参数解析失败")
		return
	}
	userId, _ := c.Get("userId")
	response, err := service.UserInfoService(request, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, response)
	return
}
