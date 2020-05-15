package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique_index"`
	Email  string `gorm:"type:varchar(100);unique_index"`
	Pwd    string
	Avatar string
	Role   int `gorm:"default: 1"` //0管理员 1普通用户
}

func QueryUserByEmailAndPwd(email, pwd string) (user User, err error) {
	logs.Info(user)
	return user, db.Where(`email = ? and pwd = ?`, email, pwd).Take(&user).Error
}
func QueryUserByName(name string) (user User, err error) {
	logs.Info("QueryUserByName", user)
	return user, db.Where(`name = ?`, name).Take(&user).Error
}
func QueryUserByEmail(email string) (user User, err error) {
	logs.Info("QueryUserByEmail", user)
	return user, db.Where(`email = ?`, email).Take(&user).Error
}
func SaveUser(user *User) error {
	logs.Info(user)
	return db.Save(user).Error
}
