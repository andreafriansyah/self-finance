package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableFinances_20210923_145157 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableFinances_20210923_145157{}
	m.Created = "20210923_145157"

	migration.Register("CreateTableFinances_20210923_145157", m)
}

// Run the migrations
func (m *CreateTableFinances_20210923_145157) Up() {
	m.SQL("CREATE TABLE finances(id serial PRIMARY KEY, type varchar(50) NOT NULL, asaltujuan VARCHAR (100) NOT NULL, jumlah bigint NOT NULL, keterangan text NOT NULL, tanggal TIMESTAMP DEFAULT CURRENT_TIMESTAMP, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT NULL)")
}

// Reverse the migrations
func (m *CreateTableFinances_20210923_145157) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE finances")
}
