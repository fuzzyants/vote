package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/vote", os.Getenv("VOTEUSER"), os.Getenv("VOTEPWD"))
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return
}
