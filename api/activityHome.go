package api

import (
	"QUZHIYOU/services/activity"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//获取首页列表

func ActivityList(that *gin.Context) {

	this:=that.Copy()

	go func() {
		time.Sleep(20*time.Second)
		service := activity.ListActivityService{}

		if err := this.ShouldBind(&service); err == nil {
			res :=  service.List()
			this.JSON(200, res)
			fmt.Println(res,"=====res======")
		} else {
			this.JSON(200, ErrorResponse(err))
		}
	}()

}


//获取活动详情信息
func ActivityInfo(this *gin.Context) {

	service:=activity.ActivityInfo{}

	if err := this.ShouldBind(&service); err == nil {
		res := service.GetActivityInfo()
		this.JSON(200, res)
	} else {
		this.JSON(200, ErrorResponse(err))
	}


}





func ActivityList1(this *gin.Context) {
	service := activity.ListActivityService{}

	if err := this.ShouldBind(&service); err == nil {
		res :=  service.List()
		this.JSON(200, res)
		fmt.Println(res,"=====res======")
	} else {
		this.JSON(200, ErrorResponse(err))
	}

}


