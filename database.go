package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() { // Open DB
	var err error
	db, err = sql.Open("sqlite3", "./kejjosdeep.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
}

type Entry struct {
	Id   int
	Name string
}

func getDBEntries() (entries []Entry) {
	rows, err := db.Query("SELECT ROWID, name FROM entries")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		e := Entry{}
		err = rows.Scan(&e.Id, &e.Name)
		entries = append(entries, e)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

func getDBEntries() (entries []Entry) {
	rows, err := db.Query("SELECT ROWID, name FROM entries")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		e := Entry{}
		err = rows.Scan(&e.Id, &e.Name)
		entries = append(entries, e)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}
