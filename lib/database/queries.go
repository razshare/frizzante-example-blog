package database

import (
	slq "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/lib/sqlc"
	"main/lib/value"
)

var Queries = sqlc.New(value.WrapFatal(slq.Open("mysql", "root:root@/forum")).Value)
