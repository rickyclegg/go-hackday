package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	devs := loadStringCollection("optimus-prime.json")

	for _, dev := range devs {
		fmt.Println("Dev: " + dev)
	}
}

func loadStringCollection(path string) (data []string) {
	jsonFile, _ := os.Open(path)

	defer jsonFile.Close()

	json.NewDecoder(jsonFile).Decode(&data)

	return data
}
