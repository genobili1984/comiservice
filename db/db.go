package dbmanager

import (
	"database/sql"
	"sync"
)

const (
	DBMaster = iota
	DBSlave
)

var masterdb *sql.DB
var slavedb *sql.DB

var once sync.Once

func openDB() {
	mdb, err := sql.Open("mysql", "root:head@5566@tcp(127.0.0.1:3306)/")
	if err == nil {
		masterdb = mdb
	}
	sdb, err := sql.Open("mysql", "root:head@5566@tcp(127.0.0.1:3306)/")
	if err == nil {
		slavedb = sdb
	}
}

func CloseDB() {
	if masterdb != nil {
		masterdb.Close()
	}

	if slavedb != nil {
		slavedb.Close()
	}
}

func GetDB(dbType int) *sql.DB {
	once.Do(func() {
		openDB()
	})
	if dbType == DBMaster {
		return masterdb
	}
	if dbType == DBSlave {
		return slavedb
	}
	return nil
}
