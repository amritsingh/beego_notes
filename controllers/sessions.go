package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type SessionsController struct {
	beego.Controller
}

func (c *SessionsController) SignupPage() {
	c.TplName = "sessions/signup.tpl"
}

func (c *SessionsController) LoginPage() {
	c.TplName = "sessions/login.tpl"
}

func (c *SessionsController) Signup() {
	email := c.GetString("email")
	password := c.GetString("password")
	confirmPassword := c.GetString("confirm_password")

	// Check if email already exists
	available := models.UserCheckAvailability(email)
	fmt.Println(available)
	if !available {
		c.Data["alert"] = "Email already exists!"
		c.TplName = "sessions/signup.tpl"
		return
	}
	if password != confirmPassword {
		c.Data["alert"] = "Password and Confirm password missmatch!"
		c.TplName = "sessions/signup.tpl"
		return
	}
	user := models.UserCreate(email, password)
	if user.Id == 0 {
		c.Data["alert"] = "Unable to create user!"
		c.TplName = "sessions/signup.tpl"
	} else {
		// Signup successful, set session
		c.SetSession("user_id", user.Id)
		c.Redirect("/notes", http.StatusSeeOther)
	}
}
