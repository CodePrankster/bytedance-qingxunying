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
	if err := c.Bind(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	code, err := service.PublishAction(request)
	if err != nil {
		common.Error(c, code)
		return
	}
	common.OK(c, code)
	return
}
