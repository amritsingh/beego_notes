package controllers

import (
	"beego_notes/models"
	"fmt"
	"net/http"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

// NotesController operations for Notes
type NotesController struct {
	beego.Controller
}

func (c *NotesController) NotesIndex() {
	user := GetUserFromSession(&c.Controller)
	// Get all notes
	notes := models.NotesGetAll(user)
	c.Data["notes"] = notes
	c.TplName = "notes/index.tpl"
}

func (c *NotesController) NotesNewForm() {
	c.TplName = "notes/new.tpl"
}

func (c *NotesController) NotesCreate() {
	name := c.GetString("name")
	content := c.GetString("content")

	models.NotesCreate(GetUserFromSession(&c.Controller), name, content)
	c.Redirect("/notes", http.StatusTemporaryRedirect)
}

func (c *NotesController) NotesShow() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(GetUserFromSession(&c.Controller), id)
	c.Data["note"] = note
	c.TplName = "notes/show.tpl"
}

func (c *NotesController) NotesEditPage() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(GetUserFromSession(&c.Controller), id)
	c.Data["note"] = note
	c.TplName = "notes/edit.tpl"
}

func (c *NotesController) NotesUpdate() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(GetUserFromSession(&c.Controller), id)
	name := c.GetString("name")
	content := c.GetString("content")
	note.Update(name, content)
	c.Redirect("/notes/"+idStr, http.StatusFound)
}

func (c *NotesController) NotesDelete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDelete(GetUserFromSession(&c.Controller), id)
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}
