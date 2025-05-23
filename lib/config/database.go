package config

import (
	sqlib "database/sql"
	"log"
)

var Database, databaseError = sqlib.Open("mysql", "root:root@/forum")

func init() {
	if nil != databaseError {
		log.Fatal(databaseError)
	}
}
