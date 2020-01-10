package serializer

import (
	"QUZHIYOU/models"
)

type Communitys struct {
	Id      uint   `json:"id"`
	Name    string `json:"cityName"`
	KeyWord string `json:"keyword"`
}

//序列化结果
type ListCommunity struct {
	Letter string     `json:"letter"`
	Data   Communitys `json:"data"`
}

func BuildCommunity(item models.Communitys) ListCommunity {
	return ListCommunity{
		Letter: item.Letter,
		Data: Communitys{
			Id:      item.ID,
			Name:    item.Name,
			KeyWord: item.KeyWord,
		},
	}
}

func BuildCommunitys(item []models.Communitys) (list []ListCommunity) {

	//定义分字母数组
	for _, v := range item {
		if v.Letter == "Y" {
			list = append(list, BuildCommunity(v))
		}
	}

	return
}
