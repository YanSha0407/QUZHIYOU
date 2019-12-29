package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//社区动态

type Diary struct {
	gorm.Model
	Name string
	Content string
	Like int
	IsLike int
	View int
	Auth string
	CommentNum int
	Address string
	Community string
	Time time.Time
}
