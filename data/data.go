package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDb(host string, port string, user string, pw string) {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/vote", user, pw, host, port)

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return
}
