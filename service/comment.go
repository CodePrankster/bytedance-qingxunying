package service

import (
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/model"
	"errors"
	"time"
)

type CommentListInfo struct {
	*common.Response
	CommentInfos []*CommentInfo `json:"comment_list"`
}
type CommentInfo struct {
	ID         uint        `json:"id"`
	User       *model.User `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"`
}

func NewCommentListInfo() *CommentListInfo {
	return &CommentListInfo{}
}

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

func (f *CommentListInfo) CommentList(request *common.CommentListRequest) (*CommentListInfo, error) {
	vid := request.VideoId
	code, commentList, err := mysql.GetCommentListByVid(vid)
	if err != nil {
		f.Response = &common.Response{
			StatusCode: code,
			StatusMsg:  common.GetMsg(code),
		}
		return nil, err
	}

	// 根据评论集合拿到对应的uid，查询出所有的用户信息
	uids := make([]uint, 0)
	for _, comment := range commentList {
		uids = append(uids, comment.Uid)
	}
	userMap, err := mysql.MQueryUserById(uids)
	if err != nil {
		return nil, err
	}
	commentInfos := make([]*CommentInfo, 0)
	for _, comment := range commentList {
		user := userMap[comment.Uid]

		commentInfos = append(commentInfos, &CommentInfo{
			ID:         comment.ID,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		})
	}
	f.CommentInfos = commentInfos
	return f, nil

}
