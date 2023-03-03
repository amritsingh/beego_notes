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
