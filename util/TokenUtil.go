package util

import (
	"dousheng-backend/common"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Id       int64  `json:"Id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetToken(userId uint, request *common.RegAndLogRequest) string {
	secretKey := []byte("bytedance")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{int64(userId), request.Username,
		request.Password, jwt.StandardClaims{}})
	tokenStr, _ := token.SignedString(secretKey)
	return tokenStr
}

func TokenVerify(tokenStr string) (uint, error) {
	//Token解析示例
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("bytedance"), nil
	})
	if err != nil {
		return 0, err
	}
	claim := token.Claims.(jwt.MapClaims)

	var parmMap map[string]interface{}

	parmMap = make(map[string]interface{})
	var userId float64
	for key, val := range claim {

		parmMap[key] = val
		if key == "Id" {
			userId = val.(float64)
		}
	}
	return uint(userId), nil

}
