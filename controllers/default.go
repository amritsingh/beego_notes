package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Notes App"
	c.Data["LoggedIn"] = (c.GetSession("user_id") != nil)
	c.TplName = "index.tpl"
}
