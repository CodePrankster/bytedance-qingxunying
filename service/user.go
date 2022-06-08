package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/model"
	"dousheng-backend/util"
	"gorm.io/gorm"
	"net/http"
)

//user_id从六位数开始向后生成
//var userIdSequence = int64(99999)

func Register(request *common.RegAndLogRequest) common.RegAndLogResponse {
	//1 长度检验
	if len(request.Username) > 32 || len(request.Password) > 32 {
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: "用户名或密码过长", UserID: -1, Token: ""}
	}
	//2 用户名重复性检验
	count := mysql.CheckUserName(request.Username)
	if count != 0 {
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: "用户名重复", UserID: -1, Token: ""}
	}
	//3 插入数据库
	registInfo := model.RegistInfo{
		Model:    gorm.Model{},
		UserName: request.Username,
		Password: request.Password,
	}
	//将信息插入注册信息表
	userId, _ := mysql.InsertRegistInfo(registInfo)
	userInfo := model.User{
		Model:    gorm.Model{ID: userId},
		Name:     request.Username,
		IsFollow: false,
	}
	//将信息插入用户信息表
	mysql.InsertUserInfo(userInfo)

	//4 生成token
	token := util.GetToken(userId, request)
	//此处AddInt64初始化问题
	//atomic.AddInt64(&userIdSequence, 1)
	return common.RegAndLogResponse{StatusCode: common.SUCCESS, StatusMsg: "注册成功", UserID: int64(userId), Token: token}
}

func Login(request *common.RegAndLogRequest) common.RegAndLogResponse {
	//1 链接数据库验证账号密码获取id
	user := mysql.SelectUser(request.Username, request.Password)
	userId := user.Model.ID
	if userId != 0 {
		//生成token
		token := util.GetToken(userId, request)
		//登录
		return common.RegAndLogResponse{StatusCode: common.SUCCESS, StatusMsg: "登录成功", UserID: int64(userId), Token: token}
	} else {
		//不存在
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: "用户不存在", UserID: -1, Token: ""}
	}
}
func UserInfo(request *common.UserInfoRequese) common.UserInfoResponse {
	//userid := c.Query("user_id")
	//token := c.Query("token")
	//fmt.Println(userid)
	//fmt.Println(token)
	//if user, exist := usersLoginInfo["zhangleidouyin"]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist!!"},
	//	})
	//}
	return common.UserInfoResponse{}
}
