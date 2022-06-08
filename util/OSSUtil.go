package util

import (
	"dousheng-backend/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Buc *oss.Bucket

func GetBuck() (error error) {
	oss.EnableCRC(false)
	//一: 读取配置文件的参数连接oss,获取oss的cli连接
	cli, err := oss.New(setting.Conf.OSSConfig.EndPoint, setting.Conf.OSSConfig.AccessKeyId, setting.Conf.OSSConfig.AccessKeySecret)
	//如果出现错误，则连接失败，记录日志直接返回
	if err != nil {
		fmt.Println("ossclient创建失败", err.Error())
		return err
	}

	//未出现错误，读取配置文件的buckeName
	exist, _ := cli.IsBucketExist(setting.Conf.OSSConfig.BucketName)
	//二: 判断bucketName是否存在，不存在需要创建，存在就直接返回即可
	if !exist {
		//bucket不存在 创建bucket
		err := cli.CreateBucket(setting.Conf.OSSConfig.BucketName)
		if err != nil {
			fmt.Println("bucket创建失败", err.Error())
			return err
		}

	}
	//bucket存在，获取bucket

	b, _ := cli.Bucket(setting.Conf.OSSConfig.BucketName)
	//将得到的bucket赋值
	Buc = b

	return nil
}
