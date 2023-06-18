package initialize

import (
	"github.com/gin-gonic/gin"
	"colly_music/controller"
)

func Route() {

	e := gin.Default()
	e.Static("/assets", "./assets")
	e.LoadHTMLGlob("tempates/*")

	e.GET("/login", controller.Login)
	e.POST("/login", controller.DoLogin)
	e.POST("/checkName", controller.CheckName)

	e.GET("/register", controller.Register)
	e.POST("/register", controller.DoRegister)
	e.POST("/doRegister", controller.DoRegisters)
	e.POST("/checkUserName", controller.CheckUserName)

	e.GET("/menu", controller.Menu)
	e.GET("/menu1", controller.Menu1)
	e.GET("/hotsong", controller.HotSong)
	e.GET("/singerlist", controller.SingerList)

	e.POST("/selectSigName", controller.SelectSigName)
	e.POST("/selectSongName", controller.SelectSongName)

	e.GET("/menu2", controller.Menu2)
	e.GET("/topshow", controller.ToShow)
	e.GET("/logout", controller.Logout)
	e.Run()
}
