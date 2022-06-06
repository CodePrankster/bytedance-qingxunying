package router

import (
	"dousheng-backend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/douyin")
	v1Favorite := v1.Group("/favorite")
	{
		v1Favorite.POST("/action/", controller.FavoriteAction)
		v1Favorite.GET("/list/", controller.FavoriteList)
	}
	v1Publish := v1.Group("/publish")
	{
		v1Publish.POST("/action/", controller.PublishAction)
		//v1Publish.GET("/list/", controller.FavoriteList)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "404",
		})
	})
	return r
}
