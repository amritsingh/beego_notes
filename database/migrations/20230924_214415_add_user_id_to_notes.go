package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddUserIdToNotes_20230924_214415 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUserIdToNotes_20230924_214415{}
	m.Created = "20230924_214415"

	migration.Register("AddUserIdToNotes_20230924_214415", m)
}

// Run the migrations
func (m *AddUserIdToNotes_20230924_214415) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notes ADD user_id bigint")
}

// Reverse the migrations
func (m *AddUserIdToNotes_20230924_214415) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE notes DROP COLUMN user_id")
}
