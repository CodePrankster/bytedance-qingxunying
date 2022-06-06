package service

import (
	"bytes"
	"context"
	"dousheng-backend/common"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/model"
	"dousheng-backend/setting"
	"dousheng-backend/util"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func PublishAction(request *common.PublishActionRequest) (int32, error) {

	//1 由于拦截器的存在，走到这里表明用户一定存在 上传用户的视频和封面图
	data := request.Data
	_, bucket, _ := util.NewOssServer().GetBuck()
	//获取用户上传文件的名称
	filename := getVideoName(data)
	//1.1上传视频并返回视频的路径
	PlayUrl := pushFile(filename, data, bucket)
	//1.2 获取封面并完成封面图的上传
	CoverUrl := getVedioFirstImg(bucket, PlayUrl, setting.Conf.FfmpegPath)
	//2 将获取的所有结果存入到数据库
	//id uid PlayUrl CoverUrl Title IsFavorite  默认false
	title := request.Title
	video := model.Video{
		Model:      gorm.Model{},
		Uid:        0,
		PlayUrl:    PlayUrl,
		CoverUrl:   CoverUrl,
		Title:      title,
		IsFavorite: false,
	}
	err := mysql.InsertVideo(video)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func pushFile(filename string, data *multipart.FileHeader, bucket *oss.Bucket) string {
	file, _ := data.Open()
	defer file.Close()
	fileByte, _ := ioutil.ReadAll(file) //获取上传文件字节流
	bucket.PutObject(filename, bytes.NewReader([]byte(fileByte)))
	return setting.Conf.OSSConfig.SufferUrl + filename
}

func getVideoName(data *multipart.FileHeader) string {
	//获取上传文件的名称
	fileName := data.Filename
	start := strings.LastIndex(fileName, ".")
	//获取文件的类型
	ext := fileName[start:]
	uid := uuid.NewV4().String()
	time := time.Now().Format("2006-01-02")
	filename := time + "/" + uid + ext
	return filename
}

func getVedioFirstImg(bucket *oss.Bucket, url string, ffmpegPath string) string {
	//1 将视频的封面图截取到本地的temp目录
	var outputerror string
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50000)*time.Millisecond)
	//ffmpeg -i http://video.pearvideo.com/head/20180301/cont-1288289-11630613.mp4 -r 1 -t 4 -f image2 image-%05d.jpeg
	coverInServer := setting.Conf.CoverInServer + uuid.NewV4().String() + ".jpg"
	cmd := exec.CommandContext(ctx, ffmpegPath,
		"-i",
		url,
		"-r",
		"1",
		"-t",
		"4",
		"-f",
		"image2",
		coverInServer)
	defer cancel()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		outputerror += fmt.Sprintf("lastframecmderr:%v;", err)
	}
	if stderr.Len() != 0 {
		outputerror += fmt.Sprintf("lastframestderr:%v;", stderr.String())
	}
	if ctx.Err() != nil {
		outputerror += fmt.Sprintf("lastframectxerr:%v;", ctx.Err())
	}

	//2 将本地的temp的封面图进行上传并返回url
	//2.1 获取用户上传文件的名称
	filename := time.Now().Format("2006-01-02") + "/" + uuid.NewV4().String() + ".jpg"
	//2.2 上传视频并返回视频的路径
	coverFile, _ := os.Open(coverInServer)
	//获取上传文件字节流
	fileByte, _ := ioutil.ReadAll(coverFile)
	bucket.PutObject(filename, bytes.NewReader([]byte(fileByte)))
	img := setting.Conf.OSSConfig.SufferUrl + filename
	return img
}
