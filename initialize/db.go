package initialize

import (
	"colly_music/global"
	"colly_music/model"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MySQLLink() {
	LoadConfig()
	m := global.Config.MySql
	fmt.Printf("m.Username: %v\n", m.Username)
	fmt.Printf("m.Password: %v\n", m.Password)
	fmt.Printf("m.Url: %v\n", m.Url)
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic(err)
		fmt.Printf("mysql error: %s\n", err)
		return
	}
	//创建表结构
	db.AutoMigrate(&model.Singer{}, &model.Song{}, &model.Comment{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Session{})
	sqlDb, err1 := db.DB()
	if err1 != nil {
		log.Panic(err1)
		fmt.Printf("mysql error: %s\n", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	global.Db = db
}
