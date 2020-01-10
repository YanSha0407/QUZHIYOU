package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeCommunityService struct {

}



func (service *HomeCommunityService)GetCommunity() serializer.Response {

	var communitys []models.Communitys

	models.PG.Find(&communitys)

	list1:=serializer.BuildCommunitys(communitys)
	return serializer.Response{
		Code:  0,
		Data: list1 ,
		Msg:   "",
		Error: "",
	}

}