package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"dousheng-backend/util"
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
	token := c.Query("token")
	userId, _ := strconv.Atoi(userIdStr)
	// 拿去当前登录用户的id
	loginUserId, _ := util.TokenVerify(token)
	if loginUserId == 0 {
		response, err := service.PublishList(uint(userId), uint(userId))
		if err != nil {
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		c.JSON(http.StatusOK, response)
	} else {
		response, err := service.PublishList(uint(userId), loginUserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		c.JSON(http.StatusOK, response)
	}

	return
}
