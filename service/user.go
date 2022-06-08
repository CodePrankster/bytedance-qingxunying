package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"dousheng-backend/util"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func Register(request *common.RegAndLogRequest) common.RegAndLogResponse {
	var errId int64 = -1
	perrId := &errId
	//1 长度检验
	if len(request.Username) > 32 || len(request.Password) > 32 {
		msg := "用户名或密码过长"
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: &msg, UserId: perrId, Token: nil}
	}
	//2 用户名重复性检验
	count := mysql.CheckUserName(request.Username)
	if count != 0 {
		msg := "用户名重复"
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: &msg, UserId: perrId, Token: nil}
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
		Model: gorm.Model{ID: userId},
		Name:  request.Username,
	}
	//将信息插入用户信息表
	mysql.InsertUserInfo(userInfo)

	//4 生成token
	token := util.GetToken(int64(userId), request)
	fmt.Println(token)

	//收集信息返回响应
	userId64 := int64(userId)
	puId := &userId64
	msg := "注册成功"
	return common.RegAndLogResponse{StatusCode: common.SUCCESS, StatusMsg: &msg, UserId: puId, Token: &token}
}

func Login(request *common.RegAndLogRequest) common.RegAndLogResponse {
	//1 链接数据库验证账号密码获取id
	user := mysql.SelectUser(request.Username, request.Password)
	userId := int64(user.Model.ID)
	var uId *int64
	uId = &userId
	//不存在
	if userId == 0 {
		msg := "用户不存在"
		return common.RegAndLogResponse{StatusCode: http.StatusBadRequest, StatusMsg: &msg, UserId: uId, Token: nil}
	}
	//生成token
	token := util.GetToken(userId, request)
	//登录
	msg := "登录成功"
	return common.RegAndLogResponse{StatusCode: common.SUCCESS, StatusMsg: &msg, UserId: uId, Token: &token}

}

func UserInfoService(request *common.UserInfoRequese, userId uint) (common.UserInfoResponse, error) {
	//查询用户的名称
	userName, err := mysql.SelectUserName(int64(userId))
	if err != nil {
		msg := "用户名称获取失败"
		return common.UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			User:       common.User{},
		}, err
	}

	//查询用户的关注总数
	followCount, err := redis.GetFollowCount(string(int64(userId)))

	if err != nil {
		msg := "用户关注总数获取失败"
		return common.UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			User:       common.User{},
		}, err
	}
	//查询用户的粉丝总数
	followerCount, err := redis.GetFollowerCount(string(int64(userId)))
	if err != nil {
		msg := "用户粉丝总数获取失败"
		return common.UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			User:       common.User{},
		}, err
	}
	//客户端查询用户是否关注
	isFollow, err := redis.IsFollow(string(int64(userId)), string(int64(userId)))
	if err != nil {
		msg := "用户粉丝总数获取失败"
		return common.UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			User:       common.User{},
		}, err
	}

	//组装用户的信息
	author := common.User{
		ID:            int64(userId),
		Name:          userName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}

	//返回用户的信息
	msg := "查询成功"
	return common.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		User:       author,
	}, nil
}
