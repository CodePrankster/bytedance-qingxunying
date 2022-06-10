// package controller

// import (
// 	"dousheng-backend/common"
// 	"dousheng-backend/service"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// // FeedList 视频feed流接口
// func FeedList(c *gin.Context) {
// 	// 参数校验
// 	latestTime := c.Query("latest_time")

// 	token := c.Query("token")
// 	request := &common.FeedRequest{
// 		LatestTime: latestTime,
// 		Token:      token,
// 	}

// 	list, _ := service.NewVideoFeedListInfo().FeedList(request)

// 	c.JSON(http.StatusOK, list)
// 	return

// }
