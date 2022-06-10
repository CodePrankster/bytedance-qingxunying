package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	userIdStr := c.Query("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	response, err := service.PublishList(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, response)
	return
}
