package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("userId")
		fmt.Println(token,"---------token----------")
	}
}
