package dao

import (
	"colly_music/global"
	"colly_music/model"
)

type QQDao struct {
}

func (q *QQDao) InsertSinger(singer *model.Singer) error {

	d := global.Db.Create(singer)
	return d.Error
}

func (q *QQDao) InsertSong(m map[string]string) error {
	song := model.Song{
		Name:        m["Name"],
		Top:         m["Top"],
		Duration:    m["Duration"],
		Vip:         m["Vip"],
		Album_name:  m["Album_name"],
		Language:    m["Language"],
		Genre:       m["Genre"],
		Issue_date:  m["Issue_date"],
		Singer_name: m["SingerName"],
	}
	d := global.Db.Create(&song)
	return d.Error
}

func (q *QQDao) InsertComment(comment [][]string) bool {
	var flag bool = true
	for _, v := range comment {
		for i := 0; i < len(v)-1; i++ {
			// fmt.Printf(" n: %v\n", v[1])
			coms := model.Comment{
				Comments: v[1],
			}
			d := global.Db.Create(&coms)
			if d.Error != nil{
				flag = false
				return flag
			}
		}
	}
	return flag
}
