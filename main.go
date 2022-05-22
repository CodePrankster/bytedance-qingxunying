package main

import (
	"dousheng-backend/router"
	"dousheng-backend/setting"
	"fmt"
)

func main() {
	// 配置初始化
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 路由初始化
	r := router.InitRouter()

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
