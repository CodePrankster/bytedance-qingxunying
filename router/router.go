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
	v1.GET("/feed/", controller.FeedList)
	v1Favorite := v1.Group("/favorite")
	{
		v1Favorite.POST("/action/", middleware.Authentication, controller.FavoriteAction)
		v1Favorite.GET("/list/", middleware.Authentication, controller.FavoriteList)
	}
	v1Publish := v1.Group("/publish")
	{

		v1Publish.POST("/action/", middleware.Authentication, controller.PublishAction)
		v1Publish.GET("/list/", middleware.Authentication, controller.PublishList)
	}
	v1User := v1.Group("/user")
	{
		v1User.POST("/login/", controller.Login)
		v1User.POST("/register/", controller.Register)
		v1User.GET("/", middleware.Authentication, controller.UserInfo)
	}

	v1Relation := v1.Group("/relation")
	{
		v1Relation.POST("/action/", middleware.Authentication, controller.RelationAvtion)
		v1Relation.GET("/follow/list/", middleware.Authentication, controller.FollowList)
		v1Relation.GET("/follower/list/", middleware.Authentication, controller.FollowerList)
	}

	v1Comment := v1.Group("/comment")
	{
		v1Comment.POST("/action/", middleware.Authentication, controller.CommentAction)
		v1Comment.GET("/list/", controller.CommentList)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "404",
		})
	})
	return r
}
