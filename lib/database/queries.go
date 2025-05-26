package database

import (
	slq "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/lib/generated"
	"main/lib/value"
)

var Queries = generated.New(value.WrapFatal(slq.Open("mysql", "root:root@/forum")).Value)
