package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func databases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	dbs := []string{}
	dbs = append(dbs, "Hello")
	dbs = append(dbs, "World")
	json.NewEncoder(w).Encode(dbs)
}

func main() {
	http.Handle("/dbs", http.HandlerFunc(databases))
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
