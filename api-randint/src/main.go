package main

import (
	"log"
	"math/rand"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/random_number", RandomNumber).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func RandomNumber(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(99)
	resp := Response{number}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type Response struct {
	Number int
}
