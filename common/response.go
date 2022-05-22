package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应类型
type Response struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, code int32, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  GetMsg(code),
		Data:       data,
	})
}
func Error(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  GetMsg(code),
		Data:       nil,
	})
}

func ErrorWithMsg(c *gin.Context, code int32, msg string) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  msg,
		Data:       nil,
	})
}
