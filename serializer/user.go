package serializer

import "QUZHIYOU/models"

type TbUser struct {
	UserId   int64  `  json:"userId" `
}

func BuildUser(user models.TbUser) *TbUser  {
	return &TbUser{UserId:user.UserId}
}