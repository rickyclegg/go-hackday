package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/api/devs", getDevs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getDevs(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("optimus-prime.json")
	dto, _ := ioutil.ReadAll(file)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dto)
}
