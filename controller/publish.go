package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	// 参数解析
	request := new(common.PublishActionRequest)
	if err := c.ShouldBind(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	userId, _ := c.Get("userId")
	code, _ := service.PublishAction(request, userId.(uint))
	common.OK(c, code)
	return
}

func PublishList(c *gin.Context) {
	// 参数解析
	request := new(common.PublishListRequest)
	if err := c.Bind(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	//code, msg, videoList := service.PublishList(request)
	return
}
