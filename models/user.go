package models

import (
	"beego_notes/helpers"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/oauth2"
)

type User struct {
	Id        uint64
	Username  string
	Password  string
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

func UserCheck(email string, password string) *User {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable(new(User)).Filter("username", email).One(&user)
	fmt.Println(err)
	fmt.Println(user)
	if user.Id == 0 {
		return nil
	}

	match := helpers.CheckPasswordHash(password, user.Password)
	fmt.Println(match)
	if match {
		return &user
	} else {
		return nil
	}
}

func UserUpdateOrCreate(token *oauth2.Token) *User {
	// Store details in OauthUser table
	oauthUser := SaveFBOauthDetails(token)

	o := orm.NewOrm()
	user := User{}
	o.QueryTable(new(User)).Filter("username", oauthUser.Email).One(&user)
	if user.Id == 0 {
		user := User{Username: oauthUser.Email,
			CreatedAt: time.Now(), UpdatedAt: time.Now()}
		o.Insert(&user)
	}
	// Update the user id in OauthUser
	oauthUser.UpdateUserID(&user)
	return (&user)
}
