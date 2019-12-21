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

	page := this.Page
	size := this.Size

	var (
		intpage, intsize int
	)

	if page == 0 {
		intpage = 1
	} else {
		intpage = page
	}

	if size == 0 {
		intsize = 10
	} else {
		intsize = size
	}

	start := (intpage - 1) * intsize


	var diary []*models.Diary

	 models.DB.
		Limit(intsize).
		Offset(start).
		Find(&diary)

	return serializer.BuildDiarys(diary)

}
