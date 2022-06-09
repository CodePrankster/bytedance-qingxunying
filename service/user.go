package service

import (
	"crypto/md5"
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"dousheng-backend/util"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//user_id从六位数开始向后生成
//var userIdSequence = int64(99999)
const secret = "byte.dance"

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
	// 密码加密
	request.Password = encryptPassword(request.Password)
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
	request.Password = encryptPassword(request.Password)
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

func UserInfoService(userId uint) (common.UserInfoResponse, error) {

	userInfo, err := GetUserBaseInfo(userId, strconv.Itoa(int(userId)))
	if err != nil {
		return common.UserInfoResponse{}, err
	}
	//返回用户的信息
	msg := "查询成功"
	return common.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		User:       userInfo,
	}, nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func GetUserBaseInfo(toUserId uint, userId string) (common.User, error) {
	//查询用户的名称
	userName, err := mysql.SelectUserName(int64(toUserId))
	if err != nil {
		return common.User{}, err
	}

	uid := strconv.Itoa(int(toUserId))
	//查询用户的关注总数
	followCount, err := redis.GetFollowCount(uid)
	if err != nil {
		return common.User{}, err
	}

	//查询用户的粉丝总数
	followerCount, err := redis.GetFollowerCount(uid)
	if err != nil {
		return common.User{}, err
	}

	//客户端查询用户是否关注
	isFollow, err := redis.IsFollow(userId, uid)
	if err != nil {
		return common.User{}, err
	}

	//组装用户的信息
	author := common.User{
		ID:            toUserId,
		Name:          userName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}
	return author, nil

}
