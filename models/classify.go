package models

type Classify struct {
	ID   int64 `gorm:"primary_key"`
	Name string
	Icon string
}
