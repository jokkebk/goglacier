package main

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
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
	Id       int            `db:"rowid"`
	ScanId   int            `db:"scan_id"`
	Filename string         `db:"filename"`
	Modified sql.NullString `db:"modified_gmt"`
	Size     sql.NullInt64  `db:"size"`
	Sha1     sql.NullString `db:"sha1"`
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
	err := db.Select(&files, `SELECT rowid, * FROM files WHERE scan_id IN
		(SELECT DISTINCT scan_id FROM scans WHERE entry_id = $1) ORDER BY scan_id, filename`, entryId)
	if err != nil {
		panic(err)
	}
	return
}
