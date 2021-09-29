package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableFinances_20210927_162525 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableFinances_20210927_162525{}
	m.Created = "20210927_162525"

	migration.Register("CreateTableFinances_20210927_162525", m)
}

// Run the migrations
func (m *CreateTableFinances_20210927_162525) Up() {
	m.SQL("CREATE TABLE finances(id serial PRIMARY KEY, type varchar(50) NOT NULL, asaltujuan VARCHAR (100) NOT NULL, jumlah bigint NOT NULL, keterangan text NOT NULL, tanggal DATE NOT NULL, is_deleted BOOLEAN DEFAULT FALSE,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT NULL)")
}

// Reverse the migrations
func (m *CreateTableFinances_20210927_162525) Down() {
	m.SQL("DROP TABLE finances")
}
