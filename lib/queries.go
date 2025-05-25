package lib

import (
	slq "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/lib/sqlc"
)

var connection, connectionError = slq.Open("mysql", "root:root@/forum")
var Queries = sqlc.New(connection)

func init() {
	if nil != connectionError {
		log.Fatal(connectionError)
	}
}
