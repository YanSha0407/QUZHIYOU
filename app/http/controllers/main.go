package controllers

import (
	"QUZHIYOU/serializer"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)


// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			fmt.Println(e)
			//field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			//tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Code: 40001,
				Msg:    "fail",
				Error:  "err",
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Code: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}