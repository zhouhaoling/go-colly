package controller

import (
	"colly_music/model"
	"colly_music/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, _ := ssion.IsLogin(c)
	if flag {
		c.HTML(200, "menu.html", gin.H{
			"username": username,
		})
	} else {
		u, _ := serce.CheckUserNameAndPassword(username, password)
		if u.ID > 0 {
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  u.Name,
				UserID:    int(u.ID),
			}
			//将Session保存到数据库中
			ssion.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			c.SetCookie(cookie.Name, cookie.Value, 60, "/", "localhost", false, true)
			//跳转页面
			// c.Redirect(301,"/menu")
			c.HTML(301, "menu.html", gin.H{
				"username": username,
				"password": password,
			})
		} else {
			c.HTML(200, "login.html", "用户不存在")
		}

	}

}


func Logout(c *gin.Context) {
	cookie, _ := c.Request.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//cookievlue 是 session.id
		ssion.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(c.Writer, cookie)
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

func DoRegisters(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := serce.InsertUser(username, password)
	if err != nil{
		c.HTML(200,"register.html",gin.H{
			"code":404,
			"message":"插入失败",
		})
	} else {
		c.HTML(200, "login.html", nil)
	}	

}
