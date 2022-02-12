package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
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
	// Windows may be missing this
	mime.AddExtensionType(".js", "application/javascript")

	http.Handle("/dbs", http.HandlerFunc(databases))
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Starting to serve GUI at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
