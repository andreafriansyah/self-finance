package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableSaldos_20210927_162529 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableSaldos_20210927_162529{}
	m.Created = "20210927_162529"

	migration.Register("CreateTableSaldos_20210927_162529", m)
}

// Run the migrations
func (m *CreateTableSaldos_20210927_162529) Up() {
	m.SQL("CREATE TABLE saldos(id serial PRIMARY KEY, nama_bank VARCHAR (20) NOT NULL, jumlah_saldo bigint NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT NULL)")
}

// Reverse the migrations
func (m *CreateTableSaldos_20210927_162529) Down() {
	m.SQL("DROP TABLE saldos")
}
