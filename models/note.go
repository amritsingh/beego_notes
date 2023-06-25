package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Note struct {
	Id        uint64
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *Note) TableName() string {
	return "notes"
}

func NotesGetAll() *[]*Note {
	o := orm.NewOrm()
	var notes []*Note
	numRows, err := o.QueryTable(new(Note)).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&notes)
	if err != nil {
		panic(err)
	}
	fmt.Println(numRows)
	return &notes
}

func NotesCreate(name string, content string) {
	o := orm.NewOrm()
	currTime := time.Now()
	note := Note{Name: name, Content: content, CreatedAt: currTime, UpdatedAt: currTime}
	id, err := o.Insert(&note)
	fmt.Println(id)
	fmt.Println(err)
}

func NotesFind(id uint64) *Note {
	o := orm.NewOrm()
	note := Note{Id: id}
	err := o.Read(&note)
	if err != nil {
		return nil
	} else {
		return &note
	}
}

func (note *Note) Update(name string, content string) {
	o := orm.NewOrm()
	note.Name = name
	note.Content = content
	o.Update(note)
}

func NotesMarkDelete(id uint64) {
	o := orm.NewOrm()
	note := NotesFind(id)
	note.DeletedAt = time.Now()
	o.Update(note)
}
