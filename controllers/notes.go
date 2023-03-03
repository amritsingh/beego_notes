package controllers

import (
	"beego_notes/models"

	beego "github.com/beego/beego/v2/server/web"
)

// NotesController operations for Notes
type NotesController struct {
	beego.Controller
}

func (c *NotesController) NotesIndex() {
	// Get all notes
	notes := models.NotesGetAll()
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}

func (c *NotesController) NotesNewForm() {
	c.TplName = "notes/new.tpl"
}
