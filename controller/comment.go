package controller

import (
	"dousheng-backend/common"
	"dousheng-backend/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentAction(c *gin.Context) {

	// 参数校验，只能支持登录的用户点赞
	//  判断用户是否登录
	uid, ok := c.Get("userId")
	if !ok {
		fmt.Println(common.GetMsg(common.CodeNeedLogin))
		common.Error(c, common.CodeNeedLogin)
		return
	}
	// 解析参数
	request := new(common.CommentRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	// 业务逻辑
	code, err := service.CommentAction(uid.(uint), request)
	if err != nil {
		common.ErrorWithMsg(c, code, err.Error())
		return
	}
	common.OK(c, code)
	return
}

func CommentList(c *gin.Context) {
	// 参数解析
	request := new(common.CommentListRequest)
	if err := c.ShouldBindQuery(request); err != nil {
		fmt.Println("参数解析失败")
		return
	}
	res, err := service.CommentList(request)
	if err != nil {
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
