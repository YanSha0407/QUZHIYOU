package main

import (
	"QUZHIYOU/models"
	"QUZHIYOU/routers"
)

func init()  {
	models.Initialized()
}

func main() {

	defer models.CloseDb()

	router:=routers.InitRouter()

	router.Run()

}



