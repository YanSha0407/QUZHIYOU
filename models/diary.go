package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//社区动态

type Diary struct {
	gorm.Model
	Avatar string
	Name       string
	Content    string
	Like       int
	IsLike     int
	View       int
	Auth       string
	CommentNum int
	Address    string
	Community  string
	Time       time.Time
}

var timeLayoutStr = "2006-01-02 15:04:05"

func (Diary *Diary) FormatTime() string {
	ts := Diary.Time.Format(timeLayoutStr) //time转string

	return ts

}
