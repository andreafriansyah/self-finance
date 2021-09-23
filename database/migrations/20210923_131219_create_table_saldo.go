package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableSaldo_20210923_131219 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableSaldo_20210923_131219{}
	m.Created = "20210923_131219"

	migration.Register("CreateTableSaldo_20210923_131219", m)
}

// Run the migrations
func (m *CreateTableSaldo_20210923_131219) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE saldo(id serial PRIMARY KEY, nama_bank VARCHAR (20) NOT NULL, jumlah_saldo bigint NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT NULL)")
}

// Reverse the migrations
func (m *CreateTableSaldo_20210923_131219) Down() {
	m.SQL("DROP TABLE saldo")
}
