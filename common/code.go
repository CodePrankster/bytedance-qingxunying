package common

// 业务状态码
const (
	SUCCESS         = 0
	ERROR           = 500
	CodePermissions = 400
	CodeNeedLogin   = 1000 + iota
	CodeInvalidToken
)

var errCode = map[int32]string{
	SUCCESS:          "OK",
	ERROR:            "Fail",
	CodeNeedLogin:    "请先登录",
	CodeInvalidToken: "无效的token",
	CodePermissions:  "无权限",
}

func GetMsg(code int32) string {
	return errCode[code]
}
