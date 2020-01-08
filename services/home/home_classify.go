package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeClassify struct {
}


func (classify *HomeClassify) GetClassify() serializer.Response {

	var classifys []models.Classify

	models.DB.Find(&classifys)

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildClassifys(classifys),
		Msg:   "",
		Error: "",
	}

}


