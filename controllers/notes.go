package controllers

import (
	"beego_notes/models"
	"net/http"

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

func (c *NotesController) NotesCreate() {
	name := c.GetString("name")
	content := c.GetString("content")

	models.NotesCreate(name, content)
	c.Redirect("/notes", http.StatusMovedPermanently)
}
