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

	l1:=serializer.BuildCommunitys(communitys)

	maps:=make(map[string]interface{})



	maps["list"]=l1
	return serializer.Response{
		Code:  0,
		Data: &maps,
		Msg:   "",
		Error: "",
	}

}