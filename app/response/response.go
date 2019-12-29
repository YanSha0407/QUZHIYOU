package response

import (
	"github.com/gin-gonic/gin"
)





func (g *Gin) Success(msg string, data interface{}) {
	g.Res(200, "success", msg, data)
}

func (g *Gin) Error(msg string, data interface{}) {
	g.Res(200, "error", msg, data)
}


type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int `json:"code"`
	Status string      `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Response conf gin.JSON
func (g *Gin) Res(httpCode int, status, msg string, data interface{}) {
	// response json
	g.C.JSON(httpCode, Response{
		Code:httpCode,
		Status: status,
		Msg:    msg,
		Data:   data,
	})
	return
}
