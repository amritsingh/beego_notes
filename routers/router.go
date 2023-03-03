package routers

import (
	"beego_notes/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/notes", &controllers.NotesController{}, "get:NotesIndex")
	beego.Router("/notes/new", &controllers.NotesController{}, "get:NotesNewForm")
	beego.Router("/notes", &controllers.NotesController{}, "post:NotesCreate")
}
