package main

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func init() { // Open DB
	var err error
	db, err = sqlx.Connect("sqlite3", "./kejjosdeep.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
}

type Entry struct {
	Id   int    `db:"rowid"`
	Name string `db:"name"`
}

type Scan struct {
	Id      int    `db:"rowid"`
	EntryId int    `db:"entry_id"`
	Date    string `db:"date_gmt"`
	Folder  string `db:"folder"`
	Skip    string `db:"skip"`
}

func dbEntries() (entries []Entry) {
	err := db.Select(&entries, "SELECT rowid, name FROM entries")
	if err != nil {
		panic(err)
	}
	return
}

func dbScans() (scans []Scan) {
	err := db.Select(&scans, "SELECT * FROM scans")
	if err != nil {
		panic(err)
	}
	return
}
