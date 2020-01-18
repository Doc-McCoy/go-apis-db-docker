package main

import (
	"fmt"
	"log"
	//"encoding/json"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "agenda"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Banco conectado!")

	router := mux.NewRouter()

	router.HandleFunc("/contato", GetAllPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPeople).Methods("GET")
	router.HandleFunc("/contato", CreatePeople).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePeople).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type Person struct {
	ID 		string 	`json:"id"`
	Nome 	string 	`json:"nome"`
	Numero	string 	`json:"numero"`
	Random 	string 	`json:"random"`
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {

}

func GetPeople(w http.ResponseWriter, r *http.Request) {

}

func CreatePeople(w http.ResponseWriter, r *http.Request) {

}

func DeletePeople(w http.ResponseWriter, r *http.Request) {

}
