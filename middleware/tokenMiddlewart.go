package middleware

import (
	"dousheng-backend/common"
	"dousheng-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication(c *gin.Context) {
	//Token解析示例

	method := c.Request.Method
	var token string
	if method == "POST" {
		token = c.PostForm("token")
	} else {
		token = c.Query("token")
	}

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("bytedance"), nil
	})
	if err != nil {
		common.Error(c, http.StatusInternalServerError)
		c.Abort()
		return
	}
	userId, _ := util.TokenVerify(token)
	c.Set("userId", userId)
	c.Next()
	return
}
