package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20230921_134238 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20230921_134238{}
	m.Created = "20230921_134238"

	migration.Register("CreateUsersTable_20230921_134238", m)
}

// Run the migrations
func (m *CreateUsersTable_20230921_134238) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users (id bigint unsigned NOT NULL AUTO_INCREMENT, username varchar(64) DEFAULT NULL, password varchar(255) DEFAULT NULL, created_at datetime(3) DEFAULT NULL, updated_at datetime(3) DEFAULT NULL, PRIMARY KEY (id))")
}

// Reverse the migrations
func (m *CreateUsersTable_20230921_134238) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}

// Run migration with this command:
// bee migrate -driver=mysql -conn="beego_notes:tmp_pwd@tcp(127.0.0.1:3306)/beego_notes?charset=utf8"
