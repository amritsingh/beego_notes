package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/oauth2"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Notes App"
	c.Data["LoggedIn"] = (c.GetSession("user_id") != nil)
	c.Data["FbAuthLink"] = FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.TplName = "index.tpl"
}
