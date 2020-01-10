package serializer

import (
	"QUZHIYOU/models"
	"fmt"
)

//序列化1
type Community struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	KeyWord string `json:"key_word"`
	Letter  string `json:"letter"`
}

//返回结果序列化
type ResComs struct {
	Letter string `json:"letter"`
	Data   []Community `json:"data"`
}


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

	var vs []Community
	var vs1 []Community

	var rs ResComs
	var rs1 ResComs

	 li:= ListRes{}

	for _, v := range item {
		if v.Letter == "Y" {
			rs.Letter = "Y"
			vs = append(vs, BuildCommunity(v))
		}

		if v.Letter == "W" {
			rs1.Letter = "W"
			vs1 = append(vs1, BuildCommunity(v))
		}
	}

	rs.Data = vs
	rs1.Data = vs1

	fmt.Println(vs, "----vs--")
	fmt.Println(rs, "----rs--")

	li.List = append(li.List, rs,rs1)


	list=li



	return

}
