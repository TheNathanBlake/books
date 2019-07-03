package main

import (
	"errors"
	"testing"
)

type mockStorage struct {
	throwDbErr bool
	returnId   int
}

func (store mockStorage) GetBook(id int) (*Book, error) {

	if store.throwDbErr {
		return &Book{}, errors.New("threw an error")
	}

	book := Book{ID: id, Title: "test", Author: "test", Publisher: "test"}

	return &book, nil
}

func (store mockStorage) InsertBook(book *Book) (int, error) {
	if store.throwDbErr {
		return -1, errors.New("threw an error")
	}
	return store.returnId, nil
}

func (store mockStorage) DeleteBook(id int) error {
	if store.throwDbErr {
		return errors.New("threw an error")
	}
	return nil
}

func (store mockStorage) UpdateBook(id int, book Book) error {
	if store.throwDbErr {
		return errors.New("threw an error")
	}
	return nil
}

func TestHandleCreate(t *testing.T) {
	serv := MainService{}
	store := mockStorage{false, 2}

	book := Book{
		Title:       "test",
		Author:      "test",
		Publisher:   "test",
		PublishDate: "04-04-2019",
		Rating:      1.5,
		Status:      2,
	}

	id, err := serv.HandleCreate(book, store)
	if err != nil {
		t.Errorf("unexpected error from service: %v", err.Error())
	}
	if id <= 0 {
		t.Errorf("ID should be greater than 0, received %d", id)
	}

	store.throwDbErr = true

	id, err = serv.HandleCreate(book, store)
	if err == nil {
		t.Error("error should be passed from storage layer")
	}
}

func TestHandleRead(t *testing.T) {
	serv := MainService{}
	store := mockStorage{false, 2}

	book, err := serv.HandleRead(2, store)
	if err != nil {
		t.Errorf("unexpected error from service: %v", err.Error())
	}
	if book.ID != 2 {
		t.Errorf("received improper book. ID: %d", book.ID)
	}

	store.throwDbErr = true
	book, err = serv.HandleRead(2, store)
	if err == nil {
		t.Error("error should be passed from storage layer")
	}
}

func TestHandleDelete(t *testing.T) {
	serv := MainService{}
	store := mockStorage{false, 2}

	err := serv.HandleDelete(2, store)
	if err != nil {
		t.Errorf("unexpected error from service: %v", err.Error())
	}
	store.throwDbErr = true
	err = serv.HandleDelete(2, store)
	if err == nil {
		t.Errorf("error should be passed from storage layer")
	}
}

func TestHandleUpdate(t *testing.T) {
	serv := MainService{}
	store := mockStorage{false, 2}

	book := Book{Title: "only updating the title"}

	err := serv.HandleUpdate(2, book, store)
	if err != nil {
		t.Errorf("unexpected error from service: %v", err.Error())
	}
	store.throwDbErr = true
	err = serv.HandleUpdate(2, book, store)
	if err == nil {
		t.Errorf("error should be passed from storage layer")
	}
}
