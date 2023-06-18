package controller

import (
	"colly_music/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingerList(c *gin.Context) {
	// var singerlist *[]model.Singer
	b, _ := ssion.IsLogin(c)
	if b != false {
		singerlist := singers.AllSinger()
		c.HTML(http.StatusOK, "singerlist.html", gin.H{
			"singerlist": singerlist,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func Menu1(c *gin.Context) {
	b, _ := ssion.IsLogin(c)
	if b != false {
		m := singers.CountSingerName()
		utils.CreateWordCloud("受欢迎的歌手", m, "likesinger.html")
		utils.CreateBarChart(m)
		c.HTML(200, "menu1.html", gin.H{
			"mList": m,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}

}

func Menu2(c *gin.Context) {
	b, _ := ssion.IsLogin(c)
	if b != false {
		mG := sons.CountSongGenre()
		utils.CreatePieChart("流派类型", mG, "genre.html")
		mL := sons.CountSongLanguage()
		utils.CreatePieChart("语言类型", mL, "language.html")
		mV := sons.CountSongVip()
		utils.CreatePieChart("歌曲级别", mV, "vip.html")
		mT := sons.CountSongIssue_Date()
		utils.CreatePieChart("时间", mT, "issuesong.html")
		c.HTML(http.StatusOK, "menu2.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func HotSong(c *gin.Context) {
	b, _ := ssion.IsLogin(c)
	if b != false {
		songlist := sons.AllSong()
		c.HTML(200, "hotsong.html", gin.H{
			"songlist": songlist,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}
}

func ToShow(c *gin.Context) {
	b, _ := ssion.IsLogin(c)
	if b != false {
		c.HTML(http.StatusOK, "topshow.html", nil)
	} else {
		c.HTML(200, "login.html", nil)
	}
}