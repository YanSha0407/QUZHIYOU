package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeCommunityService struct {

}



func (service *HomeCommunityService)GetCommunity() serializer.Response {

	var communitys models.Communitys

	models.PG.First(&communitys)

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildCommunity(communitys),
		Msg:   "",
		Error: "",
	}

}