package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)



func Getcommunity(c *gin.Context)  {


	var comunity home.HomeCommunityService
	res:=comunity.GetCommunity()
	c.IndentedJSON(200,&res)

}