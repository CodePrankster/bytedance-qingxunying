package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"fmt"
	"strconv"
)

const (
	followList = iota
	followerList
)

type UserInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserListInfo struct {
	*common.Response
	UserList []*UserInfo `json:"user_list"`
}

func NewUsesrListInfo() *UserListInfo {
	return &UserListInfo{}
}

func RelationAction(request *common.RelationActionRequest) (int32, error) {
	userId := strconv.Itoa(int(request.UserId))
	toUserId := strconv.Itoa(int(request.ToUserId))
	actionType := request.ActionType
	if actionType == 1 {
		_, err := redis.RelationActionFollow(userId, toUserId)
		if err != nil {
			return common.ERROR, err
		}
	} else {
		_, err := redis.RelationActionUnFollow(userId, toUserId)
		if err != nil {
			return common.ERROR, err
		}
	}
	return common.SUCCESS, nil
}

func (f *UserListInfo) GetUserList(request *common.FavoriteListRequest, listType interface{}) (error, *UserListInfo) {
	userId := strconv.Itoa(int(request.UserId))
	// TODO 参数校验

	var idList []string
	var relationNumber int64

	// 拿到当前用户的所有相关用户的id
	if listType == followList {
		err, arr := redis.GetFollowList(userId)
		if err != nil {
			f.Response = &common.Response{
				StatusCode: common.ERROR,
				StatusMsg:  common.GetMsg(common.ERROR),
			}
			return err, nil
		}
		n, err := redis.GetFollowCount(userId)
		if err != nil {
			f.Response = &common.Response{
				StatusCode: common.ERROR,
				StatusMsg:  common.GetMsg(common.ERROR),
			}
			return err, nil
		}
		idList = arr
		relationNumber = n
	} else if listType == followerList {
		err, arr := redis.GetFollowerList(userId)
		if err != nil {
			f.Response = &common.Response{
				StatusCode: common.ERROR,
				StatusMsg:  common.GetMsg(common.ERROR),
			}
			return err, nil
		}
		n, err := redis.GetFollowerCount(userId)
		if err != nil {
			f.Response = &common.Response{
				StatusCode: common.ERROR,
				StatusMsg:  common.GetMsg(common.ERROR),
			}
			return err, nil
		}
		idList = arr
		relationNumber = n
	} else {
		fmt.Println(1 / 0)
	}
	//根据idList查询出所有用户信息
	userList, err := mysql.MQUserListById(idList)
	if err != nil {
		f.Response = &common.Response{
			StatusCode: common.ERROR,
			StatusMsg:  common.GetMsg(common.ERROR),
		}
		return err, nil
	}

	userInfos := make([]*UserInfo, relationNumber)
	for _, user := range userList {
		userId := user.ID
		followCount, err := redis.GetFollowCount(strconv.Itoa(int(userId)))
		if err != nil {
			return err, nil
		}
		followerCount, err := redis.GetFollowerCount(strconv.Itoa(int(userId)))
		if err != nil {
			return err, nil
		}
		userInfos = append(userInfos, &UserInfo{
			ID:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      false, // TODO 查询当前用户是否关注该用户 由于当前用户的id需要从token中拿到所以暂时做不了
		})
	}
	f.UserList = userInfos
	f.Response = &common.Response{
		StatusCode: common.SUCCESS,
		StatusMsg:  common.GetMsg(common.SUCCESS),
	}
	return nil, f
}

// 获取关注列表
func (f *UserListInfo) FollowList(request *common.FavoriteListRequest) (error, *UserListInfo) {
	return f.GetUserList(request, followList)
}

// 获取粉丝列表
func (f *UserListInfo) FollowerList(request *common.FavoriteListRequest) (error, *UserListInfo) {
	return f.GetUserList(request, followerList)
}
