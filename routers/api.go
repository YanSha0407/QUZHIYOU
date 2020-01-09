package routers

import (
	"QUZHIYOU/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router := gin.Default()

	v1 := router.Group("v1")
	{
		// 使用中间价
		//v1.Use(middleware.JWTAuth())
		//v1.Use(middleware.WXToken())
		//首页活动列表
		v1.GET("/wechat/activity/selectActivityList", api.ActivityList)
		//活动详情信息
		v1.GET("/wechat/activity/selectActivicyInfoById", api.ActivityInfo)
		//获取code码
		v1.GET("/getqrcode", api.Getqrcode)
		//微信登录
		v1.GET("/wxlogin", api.WxLogin)
		//动态首页

		home := router.Group("home")

		{
			home.GET("/homediarys", api.HomeList)
			home.GET("/classify", api.Classify)
			home.GET("/ad", api.GetAd)
		}

	}

	return router

}
