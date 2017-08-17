package data

import (
	"fmt"
	"log"
	"time"
)

type Vote struct {
	Id        int64
	MaxUsers  uint32
	Slug      string
	Expires   time.Time
	Done      bool
	Options   []Option
	Ballots   []Ballot
	CreatedAt time.Time
}

func (vote *Vote) Create() (err error) {

	stmt, err := Db.Prepare("INSERT INTO Votes(maxUsers, slug, expires, done, createdAt) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(vote.MaxUsers, vote.Slug, vote.Expires, vote.Done, vote.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	} else {
		vote.Id = lastId
	}

	return
}
