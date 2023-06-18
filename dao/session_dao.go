package dao

import (
	"colly_music/global"
	"colly_music/model"

	"github.com/gin-gonic/gin"
)

type SessionDao struct {
}

// AddSession 向数据库中添加Session
func (s *SessionDao) AddSession(sess *model.Session) error {
	d := global.Db.Create(&sess)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// DeleteSession 删除数据库中的Session
func (s *SessionDao) DeleteSession(sessID string) error {
	//写sql语句
	sqlStr := "delete from session where session_id = ?"
	//执行sql
	d := global.Db.Exec(sqlStr, sessID)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

// GetSession 根据session的Id值从数据库中查询Session
func (s *SessionDao) GetSession(sessID string) (*model.Session, error) {
	//创建Session
	sess := &model.Session{}
	d := global.Db.Where("session_id = ?", sessID).First(sess)
	if d.Error != nil {
		return nil, d.Error
	}
	
	return sess, nil
}

// IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func (s *SessionDao) IsLogin(c *gin.Context) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookieValue, err := c.Cookie("user")
	if err != nil {
		return false , nil
	}
	if cookieValue != "" {
		//获取Cookie的value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := s.GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}