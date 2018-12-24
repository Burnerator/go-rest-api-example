package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Entity struct {
	ID            string         `json:"id,omitempty"`
	Firstfield    string         `json:"firstfield,omitempty"`
	Secondfield   string         `json:"secondfield,omitempty"`
	Relatedentity *Relatedentity `json:"relatedentity,omitempty"`
}
type Relatedentity struct {
	Firstfield  string `json:"firstfield,omitempty"`
	Secondfield string `json:"secondfield,omitempty"`
}

var items []Entity

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recovering items.")
	json.NewEncoder(w).Encode(items)
}

func get(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recovering item.")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Entity{})
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Creating item.")
	params := mux.Vars(r)
	var item Entity
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Deleting item.")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(items)
	}
}

// main function to boot up everything
func main() {
	items = append(items, Entity{ID: "1", Firstfield: "Value1", Secondfield: "Value 2", Relatedentity: &Relatedentity{Firstfield: "Value 1.1", Secondfield: "Value 1.2"}})

	router := mux.NewRouter()
	router.HandleFunc("/items", getAll).Methods("GET")
	router.HandleFunc("/items/{id}", get).Methods("GET")
	router.HandleFunc("/items/{id}", create).Methods("POST")
	router.HandleFunc("/items/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
