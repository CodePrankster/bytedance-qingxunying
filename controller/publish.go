package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	response, err := service.PublishList(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, response)
	return
}
