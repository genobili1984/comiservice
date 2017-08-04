package dbmanager

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBMaster = iota
	DBSlave
)

var masterdb *sql.DB
var slavedb *sql.DB

var once sync.Once

func openDB() {
	mdb, err := sql.Open("mysql", "root:head@5566@tcp(127.0.0.1:3306)/comico_test")
	if err == nil {
		masterdb = mdb
	}
	sdb, err := sql.Open("mysql", "root:head@5566@tcp(127.0.0.1:3306)/comico_test")
	if err == nil {
		slavedb = sdb
	}
}

func CloseDB() {
	if masterdb != nil {
		err := masterdb.Close()
		if err == nil {
			fmt.Println("close master db ok ")
		}
	}

	if slavedb != nil {
		err := slavedb.Close()
		if err == nil {
			fmt.Println("close slave db ok ")
		}
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
