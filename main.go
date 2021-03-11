package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/api/devs", GetDevs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetDevs(w http.ResponseWriter, _ *http.Request) {
	file, _ := os.Open("optimus-prime.json")

	w.Header().Set("Content-Type", "application/json")

	_, _ = io.Copy(w, file)
}
