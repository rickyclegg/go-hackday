package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("optimus-prime.json")

	if err != nil {
		log.Fatal(err)
	}

	devs, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/devs", GetDevs(devs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetDevs(dev []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(dev)
	}
}
