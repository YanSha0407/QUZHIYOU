package controllers

import (
	"QUZHIYOU/services/activity"
	"github.com/gin-gonic/gin"
)

//获取首页列表

func ActivityList(this *gin.Context) {
	service := activity.ListVideoService{}
	if err := this.ShouldBind(&service); err == nil {
		res := service.List()
		this.JSON(200, res)
	} else {
		this.JSON(200, ErrorResponse(err))
	}

}



//获取活动详情信息
func ActivityInfo(this *gin.Context) {

	service:=activity.ActivityInfo{}

	if err := this.ShouldBind(&service); err == nil {
		res := service.GetActivityInfo(service.ActivityId)
		this.JSON(200, res)
	} else {
		this.JSON(200, ErrorResponse(err))
	}


}






