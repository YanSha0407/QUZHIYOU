package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeClassify struct {
}

func (classify *HomeClassify) GetClassify() serializer.Response {

	var classifys []*models.Classify

	models.PG.Limit(8).Find(&classifys)


	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildClassifys(classifys),
		Msg:   "",
		Error: "",
	}

}
