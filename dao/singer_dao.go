package dao

import (
	"colly_music/global"
	"colly_music/model"
)


type SingerDao struct {
}
//查找歌手
func (s *SingerDao) FindSingName(singername string) *[]model.Singer {
	var singers []model.Singer
	d := global.Db.Table("singer").Where("name = ?", singername).Find(&singers)
	if d.Error != nil{
		return nil
	}
	return &singers
}

func (s *SingerDao) AllSinger() *[]model.Singer {
	var singers []model.Singer
	var tx []model.Singer
	d := global.Db.Distinct("name").Find(&singers)
	for _, v := range singers {
		s2 := s.FindOneSingerName(v.Name)
		tx = append(tx, *s2)
	}
	if d.Error != nil {
		return nil
	}
	return &tx
}

func (s *SingerDao) FindOneSingerName(name string) *model.Singer {
	var singer model.Singer
	d := global.Db.Where("name = ?", name).Last(&singer)
	if d.Error != nil {
		return nil
	}
	return &singer

}

// 统计歌手出现的次数
func (s *SingerDao) CountSingerName() map[string]int64 {
	var singers []model.Singer
	MSinger := make(map[string]int64)
	d := global.Db.Distinct("name").Find(&singers)
	if d.Error != nil {
		return nil
	}
	for _, v := range singers {
		var count int64
		global.Db.Table("singer").Where("name = ?", v.Name).Count(&count)
		// fmt.Printf("%v ,count: %v\n", v.Name, count)
		MSinger[v.Name] = count
	}
	return MSinger
}
