package routers

import (
	"QUZHIYOU/controllers"
	"github.com/astaxie/beego"
)




func init() {
    beego.Router("/wechat/activity/selectActivityList", &controllers.ActivityHomeListController{},"get:ActivityList")
    beego.Router("/wechat/activity/selectActivicyInfoById", &controllers.ActivityHomeListController{},"get:ActivityInfo")
}
