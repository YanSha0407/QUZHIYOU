package main

import (
	_ "QUZHIYOU/models"
	_ "QUZHIYOU/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
