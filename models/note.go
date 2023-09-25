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
	UserId    uint64
}

func (n *Note) TableName() string {
	return "notes"
}

func NotesGetAll(user *User) *[]*Note {
	o := orm.NewOrm()
	var notes []*Note
	numRows, err := o.QueryTable(new(Note)).
		Filter("user_id", user.Id).
		Filter("deleted_at__isnull", true).
		OrderBy("-updated_at").
		All(&notes)
	if err != nil {
		panic(err)
	}
	fmt.Println(numRows)
	return &notes
}

func NotesCreate(user *User, name string, content string) {
	o := orm.NewOrm()
	currTime := time.Now()
	note := Note{Name: name, Content: content, CreatedAt: currTime, UpdatedAt: currTime, UserId: user.Id}
	id, err := o.Insert(&note)
	fmt.Println(id)
	fmt.Println(err)
}

func NotesFind(user *User, id uint64) *Note {
	o := orm.NewOrm()
	note := Note{Id: id, UserId: user.Id}
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
	note.UpdatedAt = time.Now()
	o.Update(note)
}

func NotesMarkDelete(user *User, id uint64) {
	o := orm.NewOrm()
	note := NotesFind(user, id)
	note.DeletedAt = time.Now()
	o.Update(note)
}
