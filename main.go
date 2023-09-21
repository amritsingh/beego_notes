package main

import (
	"beego_notes/models"
	_ "beego_notes/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/go-sql-driver/mysql"
)

var AuthFilter = func(ctx *context.Context) {
	fmt.Println(ctx.GetCookie("user_id"))
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlDataSource, _ := web.AppConfig.String("mysqlDataSource")
	orm.RegisterDataBase("default", "mysql", mysqlDataSource)
	orm.RegisterModel(new(models.Note))
	orm.RegisterModel(new(models.User))

	web.InsertFilter("/notes/*", web.BeforeRouter, AuthFilter)
}

func main() {
	web.Run()
}
