package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/guregu/null.v4"
)

var db *sqlx.DB

func dbOpen(dbName string) { // Open DB
	var err error
	db, err = sqlx.Connect("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
}

func dbClose() {
	db.Close()
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

type File struct {
	Id       int         `db:"rowid"`
	ScanId   int         `db:"scan_id"`
	Filename string      `db:"filename"`
	Modified null.String `db:"modified_gmt"`
	Size     null.Int    `db:"size"`
	Sha1     null.String `db:"sha1"`
}

func dbEntries() (entries []Entry) {
	err := db.Select(&entries, "SELECT rowid, * FROM entries ORDER BY name")
	if err != nil {
		panic(err)
	}
	return
}

func dbScans() (scans []Scan) {
	err := db.Select(&scans, "SELECT rowid, * FROM scans ORDER BY entry_id, date_gmt")
	if err != nil {
		panic(err)
	}
	return
}

func dbFiles(entryId int) (files []File) {
	fmt.Println(entryId)
	err := db.Select(&files, `SELECT rowid, * FROM files WHERE scan_id IN
		(SELECT DISTINCT rowid FROM scans WHERE entry_id = ?) ORDER BY scan_id, filename`, entryId)
	if err != nil {
		panic(err)
	}
	return
}
