package main

import (
	"beego_notes/models"
	_ "beego_notes/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlDataSource, _ := beego.AppConfig.String("mysqlDataSource")
	orm.RegisterDataBase("default", "mysql", mysqlDataSource)
	orm.RegisterModel(new(models.Note))

}

func main() {
	beego.Run()
}
