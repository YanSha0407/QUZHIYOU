package serializer

import (
	"QUZHIYOU/models"
)

//序列化1
type Community struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	KeyWord string `json:"key_word"`
	Letter  string `json:"letter"`
}

//返回结果序列化2
type ResComs struct {
	Letter string      `json:"letter"`
	Data   []Community `json:"data"`
}

//返回结果序列化3
type ListRes struct {
	List []ResComs `json:"list"`
}

func BuildCommunity(item models.Communitys) Community {
	return Community{
		Id:      item.ID,
		Name:    item.Name,
		KeyWord: item.KeyWord,
		Letter:  item.Letter,
	}
}

func BuildCommunitys(item []models.Communitys) (list ListRes) {



	var res ResComs
	var res1 ResComs

	for _, v2 := range item {

		if v2.Letter=="A"{
			res.Data = append(res.Data, BuildCommunity(v2))
			res.Letter = v2.Letter
		}
		if v2.Letter=="B"{
			res1.Data = append(res1.Data, BuildCommunity(v2))
			res1.Letter = v2.Letter
		}

	}

	list.List = append(list.List, res,res1)

	return

}
