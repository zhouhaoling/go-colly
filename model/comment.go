package model

type Comment struct {
	Id int `gorm:"permary_key;AUTO_INCREMENT"`
	Comments string
}