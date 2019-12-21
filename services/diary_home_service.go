package services

import (
	"QUZHIYOU/app/models"
	"QUZHIYOU/serializer"
)

type DiaryHomeService struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (this *DiaryHomeService) GetAllDiary() []*serializer.Diary {

	var diary []*models.Diary

	models.DB.Find(&diary)

	return serializer.BuildDiarys(diary)

}
