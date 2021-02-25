package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonFile, _ := os.Open("optimus-prime.json")

	defer jsonFile.Close()

	var devs []string
	json.NewDecoder(jsonFile).Decode(&devs)

	for _, dev := range devs {
		fmt.Println("Dev: " + dev)
	}
}
