package main

import (
	_ "QUZHIYOU/models"
	_ "QUZHIYOU/routers"
	"github.com/astaxie/beego"
	"QUZHIYOU/services"
)

func main() {

	beego.InsertFilter("/wechat/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run()
}



