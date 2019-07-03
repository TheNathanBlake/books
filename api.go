package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

/*
BookAPI defines behavior for all RESTful requests for the books API.
*/
type BookAPI struct{}

/*
GetBook fetches books from DB data storage and returns them to the caller.
*/
func (api *BookAPI) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		msg := "ID provided must be a positive, non-zero integer value.  Received %s"
		http.Error(w, fmt.Sprintf(msg, idString), http.StatusBadRequest)
	}

	db := DBStorage{}
	serv := MainService{}

	book, err := serv.HandleRead(id, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

/*
CreateBook accepts a JSON input from the caller and stores it in the database.
*/
func (api *BookAPI) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook *Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db := DBStorage{}
	serv := MainService{}

	id, err := serv.HandleCreate(*newBook, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"id":` + strconv.Itoa(id) + `}`))
}

/*
DeleteBook removes an entry at a given ID from the database.
*/
func (api *BookAPI) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		msg := "ID provided must be a positive, non-zero integer value.  Received %s"
		http.Error(w, fmt.Sprintf(msg, idString), http.StatusBadRequest)
	}

	db := DBStorage{}
	serv := MainService{}

	err = serv.HandleDelete(id, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

}

/*
UpdateBook accepts JSON input from the caller and an ID, and updates the entry.
Note: If a field is invalid or left out of the JSON input, it will remain at its current value.
*/
func (api *BookAPI) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got to update function")
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		msg := "ID provided must be a positive, non-zero integer value.  Received %s"
		http.Error(w, fmt.Sprintf(msg, idString), http.StatusBadRequest)
	}

	var newBook *Book
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db := DBStorage{}
	serv := MainService{}

	err = serv.HandleUpdate(id, *newBook, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
