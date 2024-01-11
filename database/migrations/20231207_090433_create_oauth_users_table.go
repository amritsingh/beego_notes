package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateOauthUsersTable_20231207_090433 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateOauthUsersTable_20231207_090433{}
	m.Created = "20231207_090433"

	migration.Register("CreateOauthUsersTable_20231207_090433", m)
}

// Run the migrations
func (m *CreateOauthUsersTable_20231207_090433) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE oauth_users (id bigint NOT NULL AUTO_INCREMENT, user_id bigint NOT NULL, name varchar(128) DEFAULT NULL, provider varchar(64) NOT NULL, uuid varchar(128) NOT NULL, email varchar(128) DEFAULT NULL, access_token varchar(256) NOT NULL, expires_at datetime NOT NULL, token_type varchar(64) NOT NULL, refresh_token varchar(256) NOT NULL, profile_pic varchar(256) NOT NULL, created_at datetime NOT NULL, updated_at datetime NOT NULL, PRIMARY KEY (id), UNIQUE KEY index_oauth_users_on_provider_user_id (provider,user_id), UNIQUE KEY index_oauth_users_on_access_token (access_token), UNIQUE KEY index_oauth_users_on_provider_email (provider,email), KEY index_oauth_users_on_email (email), KEY index_oauth_users_on_user_id (user_id))")
}

// Reverse the migrations
func (m *CreateOauthUsersTable_20231207_090433) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE oauth_users")
}
