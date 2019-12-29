package routers

import (
	"QUZHIYOU/app/http/controllers"
	"QUZHIYOU/app/http/controllers/diary"
	"QUZHIYOU/app/http/controllers/qrcode"
	"QUZHIYOU/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router:=gin.Default()

	api := router.Group("v1")
	{
		// 使用中间价
		api.Use(middleware.JWTAuth())
		//活动相关的API
		api.GET("/wechat/activity/selectActivityList", controllers.ActivityList)
		api.GET("/wechat/activity/selectActivicyInfoById", controllers.ActivityInfo)
		api.GET("/homediarys", diary.HomeList)
		api.GET("/getqrcode", qrcode.Getqrcode)
	}


	return router

}

