package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableUser_20210924_155005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableUser_20210924_155005{}
	m.Created = "20210924_155005"

	migration.Register("CreateTableUser_20210924_155005", m)
}

// Run the migrations
func (m *CreateTableUser_20210924_155005) Up() {
	m.SQL("CREATE TABLE users(user_id serial PRIMARY KEY,name VARCHAR (50) NOT NULL,email VARCHAR (30) NOT NULL,gender VARCHAR (10) NOT NULL,password TEXT NOT NULL,photo VARCHAR (50) NOT NULL,token_forgot TEXT NOT NULL,token_forgot_expired_time TIMESTAMP DEFAULT NULL,is_delete BOOLEAN DEFAULT FALSE,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,created_by INT DEFAULT NULL,updated_at TIMESTAMP DEFAULT NULL,updated_by INT DEFAULT NULL)")
}

// Reverse the migrations
func (m *CreateTableUser_20210924_155005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
