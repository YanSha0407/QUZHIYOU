package serializer

import (
	"QUZHIYOU/app/models"
	"time"
)

//社区动态 序列化
type Diary struct {
	ID uint   `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
	Like int `json:"like"`
	IsLike int `json:"is_like"`
	View int `json:"view"`
	Auth string `json:"auth"`
	CommentNum int `json:"comment_num"`
	Address string `json:"address"`
	Community string `json:"community"`
	Time time.Time `json:"time"`
}

func BuildDiary(item models.Diary) Diary  {
	return Diary{
		ID:         item.ID,
		Name:       item.Name,
		Content:    item.Content,
		Like:       item.Like,
		IsLike:     item.IsLike,
		View:       item.View,
		Auth:       item.Auth,
		CommentNum: item.CommentNum,
		Address:    item.Address,
		Community:  item.Community,
		Time:       item.Time,
	}

}


func BuildDiarys(items []*models.Diary) (diarys []*Diary)  {


	for _,item:=range items{
		diary:=BuildDiary(*item)
		diarys=append(diarys,&diary)
	}

	return

}
