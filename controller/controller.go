package controller

import (
	"colly_music/dao"
	"fmt"

	"github.com/gin-gonic/gin"
)

var serce dao.UserDao
var ssion dao.SessionDao
var singers dao.SingerDao
var sons dao.SongDao

func SelectSongName(c *gin.Context) {
	songname := c.PostForm("sogname")
	songlist := sons.FindSongNames(songname)
	if songlist != nil {
		c.HTML(200, "menu.html", gin.H{
			"songlist": songlist,
		})
	} else {
		c.HTML(200, "menu.html", gin.H{
			"songlist": nil,
		})

	}

}

func CheckUserName(c *gin.Context) {
	username := c.PostForm("username")
	b := serce.FindUserName(username)
	if b == true {
		//用户名不可用
		c.Writer.Write([]byte("<font style='color:red'>用户名已存在!</font>"))
	} else {
		c.Writer.Write([]byte("<font style='color:green'>用户名可用!</font>"))
	}

}

func DoRegister(c *gin.Context) {
	DoRegisters(c)
}

func CheckName(c *gin.Context) {
	username := c.PostForm("username")
	b := serce.FindUserName(username)
	if b != true {
		c.Writer.Write([]byte("用户名不存在"))
	}
}

func SelectSigName(c *gin.Context) {

	singername := c.PostForm("signame")
	fmt.Printf("singername: %v\n", singername)

	singerlist := singers.FindSingName(singername)
	if singerlist != nil {
		c.HTML(200, "menu.html", gin.H{
			"singerlist": singerlist,
		})
	} else {
		c.HTML(200, "menu.html", gin.H{
			"singerlist": nil,
		})
	}

}

// 注册
func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

//登录

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)

}

func Menu(c *gin.Context) {
	b, s := ssion.IsLogin(c)
	if b != false {
		c.HTML(200, "menu.html", gin.H{
			"username": s.UserName,
		})
	} else {
		c.HTML(200, "login.html", nil)
	}
}
