package util

import (
	"dousheng-backend/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

type OssServer struct {
}

var once sync.Once
var ossServerInstance *OssServer
var buc *oss.Bucket

func NewOssServer() *OssServer {
	once.Do(func() {
		ossServerInstance = &OssServer{}
	})
	return ossServerInstance
}

func (that *OssServer) GetBuck() (client *oss.Client, bucket *oss.Bucket, error error) {
	//特判 如果buc不为空，则不需要走下面的流程，直接返回即可，如果为空，就按照下面的流程创建client和bucket
	if buc != nil {
		return nil, buc, nil
	}

	//一: 读取配置文件的参数连接oss,获取oss的cli连接
	cli, err := oss.New(setting.Conf.OSSConfig.EndPoint, setting.Conf.OSSConfig.AccessKeyId, setting.Conf.OSSConfig.AccessKeySecret)
	//如果出现错误，则连接失败，记录日志直接返回
	if err != nil {
		fmt.Println("ossclient创建失败", err.Error())
		return nil, nil, err
	}

	//未出现错误，读取配置文件的buckeName
	exist, _ := cli.IsBucketExist(setting.Conf.OSSConfig.BucketName)
	//二: 判断bucketName是否存在，不存在需要创建，存在就直接返回即可
	if !exist {
		//bucket不存在 创建bucket
		err := cli.CreateBucket(setting.Conf.OSSConfig.BucketName)
		if err != nil {
			fmt.Println("bucket创建失败", err.Error())
			return nil, nil, err
		}

	}
	//bucket存在，获取bucket
	b, _ := cli.Bucket(setting.Conf.OSSConfig.BucketName)
	buc = b

	//三: 将获取的bucket 返回
	return nil, buc, nil
}
