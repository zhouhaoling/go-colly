package main

import (
	"colly_music/dao"
	"colly_music/initialize"
	"time"
)

var serce dao.UserDao
var serson dao.SongDao

func main() {
	qqmusic.Colly_qqmusic()
	time.Sleep(5 * time.Second)
	qqmusic.Colly_comment()
	initialize.MySQLLink()
	initialize.Route()
}
