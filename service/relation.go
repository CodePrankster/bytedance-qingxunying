package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/redis"
	"strconv"
)

func RelationAction(request *common.RelationActionRequest, userId uint) (int32, error) {
	toUserId := strconv.Itoa(int(request.ToUserId))
	uid := strconv.Itoa(int(userId))
	actionType := request.ActionType
	if actionType == 1 {
		_, err := redis.RelationActionFollow(uid, toUserId)
		if err != nil {
			return common.ERROR, err
		}
	} else {
		_, err := redis.RelationActionUnFollow(uid, toUserId)
		if err != nil {
			return common.ERROR, err
		}
	}
	return common.SUCCESS, nil
}

func FollowList(request *common.RelationFollowListRequest) (error, common.FollowListAndFollowerListResponse) {
	userId := strconv.Itoa(int(request.UserId))
	// TODO 参数校验

	// 拿到当前用户的所有相关用户的id
	err, idList := redis.GetFollowList(userId)
	if err != nil {
		return err, common.FollowListAndFollowerListResponse{}
	}

	//根据idList查询出所有用户信息
	//var userList []common.User
	//for i := 0; i < len(idList); i++ {
	//	uid, _ := strconv.Atoi(idList[i])
	//	userList[i], _ = GetUserBaseInfo(uint(uid))
	//}
	userList := make([]common.User, 0)
	for _, id := range idList {
		toId, _ := strconv.ParseInt(id, 10, 64)
		user, _ := GetUserBaseInfo(uint(toId), userId)
		userList = append(userList, user)
	}
	msg := "查询成功"
	return nil, common.FollowListAndFollowerListResponse{
		StatusCode: "0",
		StatusMsg:  &msg,
		UserList:   userList,
	}
}

func FollowerList(request *common.RelationFollowerListRequest) (error, common.FollowListAndFollowerListResponse) {
	userId := strconv.Itoa(int(request.UserId))
	// TODO 参数校验
	// 拿到当前用户的所有相关用户的id
	err, idList := redis.GetFollowerList(userId)
	if err != nil {
		return err, common.FollowListAndFollowerListResponse{}
	}

	//根据idList查询出所有用户信息
	userList := make([]common.User, 0)
	for _, id := range idList {
		uid, _ := strconv.Atoi(id)
		user, _ := GetUserBaseInfo(uint(uid), id)
		userList = append(userList, user)
	}
	msg := "查询成功"
	return nil, common.FollowListAndFollowerListResponse{
		StatusCode: "0",
		StatusMsg:  &msg,
		UserList:   userList,
	}
}
