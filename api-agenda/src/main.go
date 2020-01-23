package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "database-pg"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "agenda"
)

var db *sql.DB
var db_err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, db_err = sql.Open("postgres", psqlInfo)
	if db_err != nil {
		panic(db_err)
	}
	defer db.Close()

	db_err = db.Ping()
	if db_err != nil {
		panic(db_err)
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

type Message struct {
	Title 	string `json:"title"`
	Message string `json:"message"`
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var persons []Person

	result, err := db.Query("SELECT id, nome, numero FROM contato")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var person Person

		err := result.Scan(&person.ID, &person.Nome, &person.Numero)
		if err != nil {
			panic(err.Error())
		}

		persons = append(persons, person)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, nome, numero FROM contato WHERE id = $1", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var person Person

	for result.Next() {
		err := result.Scan(&person.ID, &person.Nome, &person.Numero)
		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var person Person

	err := decoder.Decode(&person)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO contato (nome, numero) VALUES ($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(person.Nome, person.Numero)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	message := Message{"Info", "Contato criado com sucesso!"}
	json.NewEncoder(w).Encode(message)
}

func DeletePeople(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query, err := db.Query("DELETE FROM contato WHERE id = $1", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	w.Header().Set("Content-Type", "application/json")
	message := Message{"Info", "Contato deletado com sucesso!"}
	json.NewEncoder(w).Encode(message)
}
