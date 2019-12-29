package routers

import (
	 "QUZHIYOU/api"
	"QUZHIYOU/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router:=gin.Default()

	v1 := router.Group("v1")
	{
		// 使用中间价
		v1.Use(middleware.JWTAuth())
		//活动相关的API
		v1.GET("/wechat/activity/selectActivityList", api.ActivityList)
		v1.GET("/wechat/activity/selectActivicyInfoById", api.ActivityInfo)
		v1.GET("/homediarys", api.HomeList)
		v1.GET("/getqrcode", api.Getqrcode)
	}

	return router

}

