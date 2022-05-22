package controller

import (
	"bytes"
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	// 配置初始化
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 数据库初始化
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// redis初始化
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	m.Run()
}
func TestFavoriteAction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/douyin/favorite/action/"
	r.POST(url, FavoriteAction)
	body := `{
    "user_id":1,
    "token":"abcd",
    "video_id":1,
    "action_type":1
	}`
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	fmt.Println(w.Code)
	assert.Equal(t, 200, w.Code)
}

func TestFavoriteList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// 定义两个测试用例
	tests := []struct {
		name    string
		param   string
		wantErr bool
	}{
		{"base case", "?user_id=1&token=abc", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := "/douyin/favorite/list/"
			request, _ := http.NewRequest(http.MethodGet, url+test.param, nil)
			w := httptest.NewRecorder()
			r.GET(url, FavoriteList)
			r.ServeHTTP(w, request)
			assert.Equal(t, 200, w.Code)
		})
	}

}
