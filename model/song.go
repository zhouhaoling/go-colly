package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name string //歌曲名
	Top string  //排名
	Duration string //时长
	Vip string //是否vip
	Album_name string //专辑名
	Language string //语种
    Genre string //流派
	Issue_date string //发行时间
	Singer_name string //歌手
}