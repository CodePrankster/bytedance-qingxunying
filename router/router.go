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
	v1Favorite := v1.Group("favorite")
	{
		v1Favorite.POST("/action/", controller.FavoriteAction)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "404",
		})
	})
	return r
}
