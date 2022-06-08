package service

import (
	"context"
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/model"
	"dousheng-backend/setting"
	"dousheng-backend/util"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func PublishList(userId uint) (common.PublicListResponse, error) {
	//经过拦截器后表明请求是合法的，可以继续执行

	//通过传来的user_id查询作者和作者的视频
	//1 查询用户的信息
	//查询用户的名称
	userName, err := mysql.SelectUserName(int64(userId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "用户名称获取失败",
			VideoList:  nil,
		}, err
	}

	//查询用户的关注总数
	followCount, err := redis.GetFollowCount(string(userId))

	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "用户关注总数获取失败",
			VideoList:  nil,
		}, err
	}
	//查询用户的粉丝总数
	followerCount, err := redis.GetFollowerCount(string(userId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "用户粉丝总数获取失败",
			VideoList:  nil,
		}, err
	}
	//客户端查询用户是否关注
	isFollow, err := redis.IsFollow(string(userId), string(userId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "查询用户是否关注操作失败",
			VideoList:  nil,
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

	//2 查询视频的相关信息

	//查询视频的基础信息
	videoBaseList, err := mysql.SelectVideoListByUserId(int64(userId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "查询用户视频操作失败",
			VideoList:  nil,
		}, err
	}
	//判断用户是否有视频

	if len(videoBaseList) == 0 {
		return common.PublicListResponse{
			StatusCode: 0,
			StatusMsg:  "查询成功",
			VideoList:  nil,
		}, nil
	}
	VideoId := videoBaseList[0].ID
	PlayUrl := videoBaseList[0].PlayUrl
	CoverUrl := videoBaseList[0].CoverUrl
	Title := videoBaseList[0].Title

	//查询视频的点赞总数
	FavoriteCount, err := redis.GetVideoFavoriteNum(string(VideoId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "查询视频点赞操作失败",
			VideoList:  nil,
		}, err
	}
	//查询视频的评论总数
	CommentCount, err := mysql.GetVideoCommentNum(int64(VideoId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "查询视频评论总数操作失败",
			VideoList:  nil,
		}, err
	}
	//查询视频是否点赞
	IsFavorite, err := redis.IsFavorite(string(userId), string(VideoId))
	if err != nil {
		return common.PublicListResponse{
			StatusCode: 1,
			StatusMsg:  "查询视频是否点赞操作失败",
			VideoList:  nil,
		}, err
	}
	//3 封装数据返回

	videos := make([]common.Video, len(videoBaseList))

	for i := 0; i < len(videos); i++ {
		videos[i].Id = int64(VideoId)
		videos[i].Author = author
		videos[i].PlayUrl = PlayUrl
		videos[i].CoverUrl = CoverUrl
		videos[i].FavoriteCount = FavoriteCount
		videos[i].CommentCount = CommentCount
		videos[i].IsFavorite = IsFavorite
		videos[i].Title = Title
	}

	return common.PublicListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  videos,
	}, nil
}

func PublishAction(request *common.PublishActionRequest, userId uint) (int32, error) {
	//用于防止视频重复
	uid := uuid.NewV4().String()
	//用于oss的二级目录
	timeStr := time.Now().Format("2006-01-02")
	//获取封面图在oss的路径
	CoverUrl := setting.Conf.OSSConfig.SufferUrl + strconv.Itoa(int(userId)) + "/" + timeStr + "/" + uid + ".jpg"

	//1 由于拦截器的存在，走到这里表明用户一定存在 上传用户的视频和封面图
	//获取用户上传视频的名称
	filename := getVideoName(uid, userId, request.Data)
	//1.1上传视频并返回视频的路径
	PlayUrl := pushFile(filename, request.Data, util.Buc)

	//1.2 使用PlayUrl完成完成封面图的上传
	go getVedioFirstImg(timeStr, uid, userId, util.Buc, PlayUrl, setting.Conf.FfmpegPath)

	//2 将获取的所有结果存入到数据库
	video := model.Video{
		Model:    gorm.Model{},
		Uid:      userId,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
		Title:    request.Title,
	}

	err := mysql.InsertVideo(&video)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return common.SUCCESS, nil
}

func getVideoName(uid string, userId uint, data *multipart.FileHeader) string {
	//获取上传文件的名称
	fileName := data.Filename
	start := strings.LastIndex(fileName, ".")
	//获取文件的类型
	ext := fileName[start:]
	filename := strconv.Itoa(int(userId)) + "/" + time.Now().Format("2006-01-02") + "/" + uid + ext
	return filename
}

func pushFile(filename string, data *multipart.FileHeader, bucket *oss.Bucket) string {
	file, _ := data.Open()

	defer file.Close()
	bucket.PutObject(filename, file)

	return setting.Conf.OSSConfig.SufferUrl + filename
}

func getVedioFirstImg(timeStr string, uid string, userId uint, bucket *oss.Bucket, url string, ffmpegPath string) {
	//temp文件
	coverInServer := setting.Conf.CoverInServer + uid + ".jpg"
	//1 将视频的封面图截取到本地的temp目录
	var outputerror string
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50000)*time.Millisecond)
	//ffmpeg -i http://video.pearvideo.com/head/20180301/cont-1288289-11630613.mp4 -r 1 -t 4 -f image2 image-%05d.jpeg
	cmd := exec.CommandContext(ctx, ffmpegPath,
		"-i",
		url,
		"-ss",
		"1",
		"-f",
		"image2",
		coverInServer)

	defer cancel()
	err := cmd.Run()
	if err != nil {
		outputerror += fmt.Sprintf("lastframecmderr:%v;", err)
	}
	//2 将本地的temp的封面图进行上传
	filename := strconv.Itoa(int(userId)) + "/" + timeStr + "/" + uid + ".jpg"
	bucket.PutObjectFromFile(filename, coverInServer)
}
