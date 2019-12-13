package routers

import (
	"QUZHIYOU/app/http/controllers/activity"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router:=gin.Default()


	//活动相关的API
	router.GET("/wechat/activity/selectActivityList", activity.ActivityList)
	router.GET("/wechat/activity/selectActivicyInfoById", activity.ActivityInfo)


	return router

}

