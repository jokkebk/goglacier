package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
)

func printJson(data interface{}) {
	bs, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bs))
}

func entries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	es := dbEntries()
	json.NewEncoder(w).Encode(es)
	printJson(es)
}

func scans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	ss := dbScans()
	json.NewEncoder(w).Encode(ss)
	printJson(ss)
}

func files(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fs := dbFiles(1)
	json.NewEncoder(w).Encode(fs)
	printJson(fs)
}

func scan(w http.ResponseWriter, r *http.Request) {
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

	http.Handle("/entries", http.HandlerFunc(entries))
	http.Handle("/scans", http.HandlerFunc(scans))
	http.Handle("/scan", http.HandlerFunc(scan))
	http.Handle("/files", http.HandlerFunc(files))
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Starting to serve GUI at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
