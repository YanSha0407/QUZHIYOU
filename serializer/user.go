package serializer

import "QUZHIYOU/models"

type User struct {
	Id uint `  json:"id" `
}

func BuildUser(user *models.User) *User {
	return &User{Id: user.ID}
}
