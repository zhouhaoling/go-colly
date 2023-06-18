package dao

import (
	"colly_music/global"
	"colly_music/model"
)

type SongDao struct {
}

//
func (s *SongDao) FindSongNames(songname string) *[]model.Song {
	var songs []model.Song
	d := global.Db.Table("song").Where("name = ?", songname).Find(&songs)
	if d.Error != nil {
		return nil
	} 
	return &songs
}

// 查询指定歌曲名的数据(返回一个)
func (s *SongDao) FindSongName(songname string) *model.Song {
	song := model.Song{}
	d := global.Db.Where("name = ?", songname).Last(&song)
	if d.Error != nil {
		return nil
	}
	return &song
}

// 查找全部歌曲数据
func (s *SongDao) AllSong() *[]model.Song {
	var singers []model.Song
	var tx []model.Song
	d := global.Db.Distinct("name").Find(&singers)
	for _, v := range singers {
		s2 := s.FindSongName(v.Name)
		tx = append(tx, *s2)
	/* 	if s2 != nil {
			tx = append(tx, *s2)
		}else{
			continue 
		}*/
	}
	if d.Error != nil {
		return nil
	}
	return &tx
}

//统计vip
func(s *SongDao) CountSongVip() map[string]int64 {
	var songs []model.Song
	Mvip := make(map[string]int64)
	d := global.Db.Distinct("vip").Find(&songs)
	if d.Error != nil {
		return nil
	}
	for _, v := range songs {
		// fmt.Printf("v: %v\n", v)
		var count int64
		global.Db.Table("song").Where("vip = ?", v.Vip).Count(&count)
		// fmt.Printf("%v ,count: %v\n", v.Genre,count)
		Mvip[v.Vip] = count
	}
	return Mvip
}

// 统计歌曲流派
func (s *SongDao) CountSongGenre() map[string]int64 {
	var songs []model.Song
	Mgenre := make(map[string]int64)
	d := global.Db.Distinct("genre").Find(&songs)
	if d.Error != nil {
		return nil
	}
	for _, v := range songs {
		// fmt.Printf("v: %v\n", v)
		var count int64
		global.Db.Table("song").Where("genre = ?", v.Genre).Count(&count)
		// fmt.Printf("%v ,count: %v\n", v.Genre,count)
		Mgenre[v.Genre] = count
	}
	return Mgenre
}

// 统计歌曲语种
func (s *SongDao) CountSongLanguage() map[string]int64 {
	var songs []model.Song
	MLanguage := make(map[string]int64)
	d := global.Db.Distinct("language").Find(&songs)
	if d.Error != nil {
		return nil
	}
	for _, v := range songs {
		var count int64
		global.Db.Table("song").Where("language = ?", v.Language).Count(&count)
		// fmt.Printf("%v ,count: %v\n", v.Language, count)
		MLanguage[v.Language] = count
	}
	return MLanguage
}

func (s *SongDao) CountSongIssue_Date() map[string]int64 {
	Mt := make(map[string]int64)
	var count int64
	global.Db.Table("song").Where("issue_date BETWEEN ? AND ?", "2020", "2023-12-31").Count(&count)
	Mt["2020-2023"] = count
	global.Db.Table("song").Where("issue_date BETWEEN ? AND ?", "2016", "2019-12-31").Count(&count)
	Mt["2016-2019"] = count
	global.Db.Table("song").Where("issue_date BETWEEN ? AND ?", "2012", "2015-12-31").Count(&count)
	Mt["2012-2015"] = count
	global.Db.Table("song").Where("issue_date BETWEEN ? AND ?", "1990", "2011-12-31").Count(&count)
	Mt["1990-2011"] = count
	return Mt
}
