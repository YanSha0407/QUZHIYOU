package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type AddDiaryService struct {
	Avatar    string `form:"avatar" json:"avatar" `
	Name      string `form:"name" json:"name" `
	Content   string `form:"content" json:"content" `
	Address   string `form:"address" json:"address" `
	Community string `form:"community" json:"community"`
	Photos []string `form:"photos" json:"photos"`
	Tag string `form:"tag" json:"tag"`
	PhotosThumb []string `form:"photosthumb" json:"photosthumb"`
}

func (diary *AddDiaryService) AddDiary() serializer.Response {

	dia := models.Diary{
		Avatar:    diary.Avatar,
		Name:      diary.Name,
		Content:   diary.Content,
		Address:   diary.Address,
		Community: diary.Community,
		Photos:diary.Photos,
		PhotosThumb:diary.Photos,
		Tag:diary.Tag,
	}

	models.PG.Create(&dia)

	return serializer.Response{
		Code:  0,
		Data:  "创建成功",
		Msg:   "",
		Error: "",
	}

}
