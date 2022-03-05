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

func Entries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	es := dbEntries()
	json.NewEncoder(w).Encode(es)
	//printJson(es)
}

func Scans(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	ss := dbScans()
	json.NewEncoder(w).Encode(ss)
	//printJson(ss)
}

func Files(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if scanId, err := strconv.Atoi(ps.ByName("scanid")); err != nil {
		fmt.Println("Invalid scan id")
	} else {
		fs := dbFiles(scanId)
		json.NewEncoder(w).Encode(fs)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Oh hi!")
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
	router.GET("/entries", Entries)
	router.GET("/scans", Scans)
	router.GET("/scan", DoScan)
	router.GET("/files/:scanid", Files)
	router.NotFound = http.FileServer(http.Dir("static"))
	//router.ServeFiles("/*filepath", http.Dir("static"))

	fmt.Println("Starting to serve GUI at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
