package api

import (
	"QUZHIYOU/services/wxLogin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WxLogin(c *gin.Context)  {

	code:=wxLogin.User{}

	if err:=c.ShouldBind(&code);err!=nil{
		fmt.Println("code 失败")
	}

	res:=code.WxUserLogin()


	c.JSON(200,&res)

}
