package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {
	http.HandleFunc("/encode", en)
	http.HandleFunc("/decode", de)
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

func de(w http.ResponseWriter, req *http.Request) {
	var p1 person
	if err := json.NewDecoder(req.Body).Decode(&p1); err != nil {
		log.Println("Decoded bad data", err)
	}
	fmt.Println("Person: ", p1)
}

// command to test decoding json
// curl -XGET -H "Content-type: application/json" -d '{"Name":"Siri Tali"}' 'localhost:8080/decode'
