package data

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Vote struct {
	Id        int64
	Title     string
	MaxUsers  uint32
	Slug      string
	Expires   time.Time
	Done      bool
	Options   []Option
	Ballots   []Ballot
	CreatedAt time.Time
}

// Save generates a slug, persists the struct to the db,
// and sets the corresponding Id field
func (vote *Vote) Save() (slug string, err error) {

	// TODO: use UPDATE if Id field is already set
	stmt, err := Db.Prepare(
		"INSERT INTO Votes(maxUsers, slug, expires, done, createdAt) VALUES(?, ?, ?, ?, ?)")

	if err != nil {
		return "", err
		log.Fatal("Error creating prepared statement: ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		vote.MaxUsers, vote.Slug, vote.Expires, vote.Done, vote.CreatedAt)

	if err != nil {
		return "", err
		log.Fatal("Error executing prepared statement: ", err)
	}
	fmt.Println(res)

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error getting id: ", err)
		return "", err
	} else {
		vote.Id = lastId
		// TODO: generate a short slug
		// let's use the id in hex for now
		vote.Slug = strconv.FormatInt(lastId, 16)
	}

	return vote.Slug, nil
}
