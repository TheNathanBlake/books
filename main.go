package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	bookAPI := BookAPI{}

	r.HandleFunc("/book", bookAPI.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", bookAPI.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", bookAPI.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", bookAPI.DeleteBook).Methods("DELETE")

	err := DBStorage{}.VerifyTables()
	if err != nil {
		log.Panic(err.Error())
	}

	log.Print(http.ListenAndServe(":8080", r))
}
