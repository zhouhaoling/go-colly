package dao

import (
	"colly_music/global"
	"colly_music/model"
	"fmt"
)

//用户登录注册查询
type UserDao struct {
}

func (u *UserDao) InsertUser(name string, password string)  error {
	users := model.User{
		Name: name,
		Password: password,
		Status: 1,
	}
	d2 := global.Db.Create(&users)
	
	return  d2.Error
}

func (u *UserDao) FindUserName(name string) bool {
	user := model.User{}
	d := global.Db.Where("name = ?", name).First(&user)
	if d.Error != nil {
		fmt.Println("用户不存在!")
		return false
	}
	fmt.Println("用户已存在!")
	return true
}



func (u *UserDao) CheckUserNameAndPassword(username string, password string) (*model.User, error){
	user := model.User{}
	d := global.Db.Where("name = ? AND password = ?", username, password).First(&user)
	fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)
	fmt.Printf("user.ID: %v user.Name: %v user.Password:%v\n", user.ID, user.Name,user.Password)
	if d.Error != nil {
		return &user, d.Error
	} 
	return &user, nil
}
