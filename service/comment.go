package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/model"
	"errors"
	"strconv"
	"time"
)

func CommentAction(uid uint, request *common.CommentRequest) (int32, error) {
	value := request.ActionType
	if value == 1 {
		// 说明是发布评论
		comment := &model.Comment{
			//  Uid需要通过上下文获取
			Uid:        uid,
			Vid:        request.VideoId,
			Content:    request.CommentText,
			CreateDate: time.Now().String(),
		}
		return mysql.InsertComment(comment)
	}
	if value == 2 {
		// 说明是删除评论
		return mysql.DeleteComment(request.CommentId)
	}
	return common.ERROR, errors.New("参数传入错误")
}

func CommentList(request *common.CommentListRequest) (common.CommentListResponse, error) {
	vid := request.VideoId
	code, commentList, err := mysql.GetCommentListByVid(vid)
	if err != nil {
		msg := "评论列表查询失败"
		return common.CommentListResponse{
			StatusCode:  code,
			StatusMsg:   &msg,
			CommentList: nil,
		}, err
	}

	comments := make([]common.Comment, 0)
	for _, comment := range commentList {
		user, _ := GetUserBaseInfo(comment.Uid, strconv.Itoa(int(comment.Uid)))
		comments = append(comments, common.Comment{
			ID:         int64(comment.ID),
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		})
	}
	msg := "查询评论操作成功"
	return common.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   &msg,
		CommentList: comments,
	}, nil

}
