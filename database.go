package main

import (
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

type Scan struct {
	Id     int    `db:"rowid"`
	Entry  string `db:"entry"`
	Date   string `db:"date_gmt"`
	Folder string `db:"folder"`
	Skip   string `db:"skip"`
}

type File struct {
	Id       int         `db:"rowid"`
	ScanId   int         `db:"scan_id"`
	Op       int         `db:"op"`
	Filename string      `db:"filename"`
	Modified null.String `db:"modified_gmt"`
	Size     null.Int    `db:"size"`
	Sha1     null.String `db:"sha1"`
}

func dbScans() (scans []Scan) {
	err := db.Select(&scans, "SELECT rowid, * FROM scans ORDER BY entry, date_gmt")
	if err != nil {
		panic(err)
	}
	return
}

func dbFiles(scanId int) (files []File) {
	err := db.Select(&files, `SELECT rowid, * FROM files WHERE scan_id = ?`, scanId)
	if err != nil {
		panic(err)
	}
	return
}

func dbFilecount(scanId int) (count int) {
	err := db.Get(&count, `SELECT COUNT(1) FROM files WHERE scan_id = ?`, scanId)
	if err != nil {
		panic(err)
	}
	return
}
