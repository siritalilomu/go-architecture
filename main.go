package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", en)
	http.ListenAndServe(":8080", nil)
}

func en(w http.ResponseWriter, req *http.Request) {
	p1 := person{
		Name: "Siri",
	}
	if err := json.NewEncoder(w).Encode(p1); err != nil {
		log.Println("Encoded bad data", err)
	}
}
