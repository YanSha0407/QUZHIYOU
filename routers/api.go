package routers

import (
	"QUZHIYOU/app/http/controllers/activity"
	"QUZHIYOU/app/http/controllers/diary"
	"QUZHIYOU/app/http/controllers/qrcode"
	"github.com/gin-gonic/gin"
	"QUZHIYOU/app/http/middleware"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router:=gin.Default()

	api := router.Group("v1")
	{
		// 使用中间价
		api.Use(middleware.JWTAuth())
		//活动相关的API
		api.GET("/wechat/activity/selectActivityList", activity.ActivityList)
		api.GET("/wechat/activity/selectActivicyInfoById", activity.ActivityInfo)
		api.GET("/homediarys", diary.HomeList)
		api.GET("/getqrcode", qrcode.Getqrcode)
	}


	return router

}

