package routers

import (
	"QUZHIYOU/app/http/controllers/activity"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router:=gin.Default()

	api := router.Group("v1")
	{
		//活动相关的API
		api.GET("/wechat/activity/selectActivityList", activity.ActivityList)
		api.GET("/wechat/activity/selectActivicyInfoById", activity.ActivityInfo)
	}



	return router

}

