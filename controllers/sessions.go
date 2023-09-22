package controllers

import (
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
