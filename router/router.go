package router

import (
	"dousheng-backend/controller"
	"dousheng-backend/middleware"
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
		v1Publish.POST("/action/", middleware.Authentication, controller.PublishAction)
		//v1Publish.GET("/list/", controller.FavoriteList)
	}
	v1User := v1.Group("/user")
	{
		v1User.POST("/login/", controller.Login)
		v1User.POST("/register/", controller.Register)
		//v1Publish.GET("/list/", controller.FavoriteList)
	}

	v1Relation := v1.Group("relation")
	{
		v1Relation.POST("/action/", controller.RelationAvtion)
		v1Relation.GET("/follow/list/", controller.FollowList)
		v1Relation.GET("/follower/list/", controller.FollowerList)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "404",
		})
	})
	return r
}
