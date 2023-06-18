package model

import "gorm.io/gorm"

type Singer struct {
	gorm.Model
	Name      string //作者姓名
	Singles   int //单曲数
	Albums    int //专辑数
	Mvs       int //Mv数量
	Attention float32 //关注量
	Content   string //简介
}