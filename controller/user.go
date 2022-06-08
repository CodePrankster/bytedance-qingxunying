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

//待开发
func UserInfo(c *gin.Context) common.UserInfoResponse {
	// 参数解析
	request := new(common.UserInfoRequese)
	err := c.Bind(request)
	if err != nil {
		return common.UserInfoResponse{StatusCode: http.StatusBadRequest, StatusMsg: "参数解析错误", User: common.User{}}
	}
	response := service.UserInfo(request)
	return response
}
