package main

import (
	"fmt"
)

/*
Service acts as a middle-man between the storage and API layers, which opens up
the logic for database-independent testing.
*/
type Service interface {
	HandleCreate(book *Book) (int, error)
	HandleRead(bookID int, store Storage) (*Book, error)
	HandleUpdate(bookID int, book Book, store Storage) error
	HandleDelete(bookID int, store Storage) error
}

/*
Main implementation of the Service interface.
*/
type MainService struct {
}

/*
HandleCreate validates input from user before passing the new book to the storage layer.
*/
func (serv *MainService) HandleCreate(book Book, store Storage) (int, error) {
	if !book.IsValidTitle() {
		return -1, fmt.Errorf("invalid title.  Expected non-empty string, received %s", book.Title)
	} else if !book.IsValidAuthor() {
		return -1, fmt.Errorf("invalid author.  Expected non-empty string, received %s", book.Author)
	} else if !book.IsValidPublisher() {
		return -1, fmt.Errorf("invalid publisher.  Expected non-empty string, received %s", book.Publisher)
	} else if !book.IsValidPublishDate() {
		return -1, fmt.Errorf("invalid publish date.  Expected date formatted MM-DD-YYYY, received %s", book.PublishDate)
	} else if !book.IsValidRating() {
		return -1, fmt.Errorf("invalid rating.  Expected value between 1.0-3.0, received %f", book.Rating)
	}

	id, err := store.InsertBook(&book)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (serv *MainService) HandleRead(id int, store Storage) (*Book, error) {
	book, err := store.GetBook(id)
	if err != nil {
		return &Book{}, err
	}
	return book, err
}

func (serv *MainService) HandleDelete(id int, store Storage) error {
	err := store.DeleteBook(id)
	return err
}

func (serv *MainService) HandleUpdate(id int, book Book, store Storage) error {
	err := store.UpdateBook(id, book)
	return err
}
