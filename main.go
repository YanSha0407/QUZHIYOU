package main

import (
	"QUZHIYOU/models"
	"QUZHIYOU/routers"
	"github.com/joho/godotenv"
)

func init()  {
	godotenv.Load()
	models.Initialized()
}

func main() {
	defer models.CloseDb()
	router:=routers.InitRouter()
	router.Run(":8080")
}



