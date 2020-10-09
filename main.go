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

func en(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Name: "Siri",
	}
	p2 := person{
		Name: "Keaton",
	}
	people := []person{p1, p2}
	if err := json.NewEncoder(w).Encode(people); err != nil {
		log.Println("Encoded bad data", err)
	}
}

func de(w http.ResponseWriter, r *http.Request) {
	people := []person{}
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		log.Println("Decoded bad data", err)
	}
	fmt.Println("People: ", people)
}

// command to test decoding json
// curl -XGET -H "Content-type: application/json" -d '[{"Name":"Siri Tali"}]' 'localhost:8080/decode'
