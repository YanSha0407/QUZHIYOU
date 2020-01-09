package serializer

import "QUZHIYOU/models"

type User struct {
	Id uint `  json:"userId" `
}

func BuildUser(user *models.User) *User {
	return &User{Id: user.ID}
}
