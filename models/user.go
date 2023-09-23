package models

import (
	"beego_notes/helpers"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func UserCheckAvailability(email string) bool {
	o := orm.NewOrm()
	var user User
	fmt.Println(user.TableName())
	o.QueryTable(new(User)).Filter("username", email).One(&user)
	return (user.Id == 0) // if ID == 0, email is not signed up, hence available
}

func UserCreate(email string, password string) *User {
	o := orm.NewOrm()
	hshPasswd, _ := helpers.HashPassword(password)
	entry := User{Username: email, Password: hshPasswd}
	o.Insert(&entry)
	return &entry
}

func UserFind(id uint64) *User {
	var user User
	o := orm.NewOrm()
	o.QueryTable(new(User)).Filter("id", id).One(&user)
	return &user
}
