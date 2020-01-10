package models

import "github.com/jinzhu/gorm"

type Community struct {
	gorm.Model
	Name string

}