package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name          string
	Port          int
	CoverInServer string `mapstructure:"cover_in_server"`
	FfmpegPath    string `mapstructure:"ffmpeg_path"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*OSSConfig    `mapstructure:"oss"`
}

// 定义mysql配置文件的结构体
type MySQLConfig struct {
	Host         string `mapstructure:"host"  `
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// 定义redis配置文件的结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// 定义oss配置文件的结构体
type OSSConfig struct {
	EndPoint        string `mapstructure:"end_point"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
	SufferUrl       string `mapstructure:"suffer_url"`
}

func Init() error {
	//viper.SetConfigFile("./conf/config.yml")   为了测试通过 先注释掉你
	viper.SetConfigFile("E:\\program_WorkSpace\\school\\Projects\\byteDance-project\\dousheng-backend\\conf\\config.yml")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}
