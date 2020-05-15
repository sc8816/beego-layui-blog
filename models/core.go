package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	logs.Info("初始化")
	var err error
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed" + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	var count int
	if err := db.Model(&User{}).Count(&count).Error; count == 0 && err == nil {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@qq.com",
			Pwd:    "123456",
			Avatar: "/static/images/user.jpg",
			Role:   0,
		})
	}
}
