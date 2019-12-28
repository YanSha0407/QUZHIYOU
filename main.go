package main

import (
	"QUZHIYOU/app/models"
	"QUZHIYOU/routers"
)

func init()  {
	models.Initialized()
}

func main() {

	defer models.CloseDb()

	router:=routers.InitRouter()
	router.Run(":8080")

	//router.RunTLS(":8080","./server.pem","./server.key")

}



