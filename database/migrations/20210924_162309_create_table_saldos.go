package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableSaldos_20210924_162309 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableSaldos_20210924_162309{}
	m.Created = "20210924_162309"

	migration.Register("CreateTableSaldos_20210924_162309", m)
}

// Run the migrations
func (m *CreateTableSaldos_20210924_162309) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE saldos(id serial PRIMARY KEY, nama_bank VARCHAR (20) NOT NULL, jumlah_saldo bigint NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT NULL)")

}

// Reverse the migrations
func (m *CreateTableSaldos_20210924_162309) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE saldos")
}
