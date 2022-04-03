package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func printJson(data interface{}) {
	bs, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bs))
}

func Scans(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dbScans())
}

func Files(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if scanId, err := strconv.Atoi(ps.ByName("scanid")); err != nil {
		fmt.Println("Invalid scan id")
	} else {
		json.NewEncoder(w).Encode(dbFiles(scanId))
	}
}

func Filecount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if scanId, err := strconv.Atoi(ps.ByName("scanid")); err != nil {
		fmt.Println("Invalid scan id")
	} else {
		io.WriteString(w, strconv.Itoa(dbFilecount(scanId)))
	}
}

func DoScan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var s Scan
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return // Silently fail on invalid POST data
	}
	fmt.Println(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Oh hi!")
}

// Middleware for a standard handler returning a "github.com/julienschmidt/httprouter" Handle
func middleCORS(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing SQLite database as parameter!")
	}
	dbOpen(os.Args[1])
	defer dbClose()

	// Windows may be missing this
	mime.AddExtensionType(".js", "application/javascript")

	router := httprouter.New()
	router.GET("/scans", middleCORS(Scans))
	router.GET("/scan", middleCORS(DoScan))
	router.GET("/files/:scanid", middleCORS(Files))
	router.GET("/filecount/:scanid", middleCORS(Filecount))
	router.NotFound = http.FileServer(http.Dir("static"))

	fmt.Println("Starting to serve GUI at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
