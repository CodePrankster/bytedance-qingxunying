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

	header := c.GetHeader("Content-Type")
	var token string
	if header == "" { //表明token通过参数传递
		token = c.Query("token")
	} else { //表明token通过body传递
		token = c.PostForm("token")
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
