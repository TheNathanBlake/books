package main

import (
	"database/sql"
	"fmt"
	"log"
)

/*
Storage defines methods for interacting with book storage (libraries)
*/
type Storage interface {
	InsertBook(book *Book) (int, error)
	UpdateBook(id int, book Book) error
	DeleteBook(id int) error
	GetBook(id int) (*Book, error)
}

const (
	host     = "localhost"
	port     = 15432
	user     = "postgres"
	password = "dopesauce"
	database = "postgres"
)

/*
VerifyTables verifies that all tables exist in database
*/
func (storage DBStorage) VerifyTables() error {
	pg, err := sql.Open("postgres", getConnectionInfo())
	if err != nil {
		return err
	}
	defer pg.Close()

	dbCreate := `
	create table if not exists book_status (
		status_id numeric(1) not null primary key,
		status_name varchar not null
	);
	
	insert into book_status (status_id, status_name) 
	values (1, 'CheckedIn') 
	on conflict (status_id) do nothing;
	
	insert into book_status (status_id, status_name) 
	values (2, 'CheckedOut') 
	on conflict (status_id) do nothing;
	
	insert into book_status (status_id, status_name) 
	values (3, 'OnHold') 
	on conflict (status_id) do nothing;
	
	create table if not exists book (
		id serial primary key,
		title varchar not null,
		author varchar not null,
		publisher varchar not null,
		publish_date date,
		rating numeric(6, 5) not null,
		status numeric not null,
		foreign key (status) references book_status(status_id),
		unique (title, author, publisher)
	);
	`
	_, err = pg.Exec(dbCreate)
	return err
}

/*
DBStorage acts as database storage
*/
type DBStorage struct{}

/*
InsertBook communicates with the Postgres instance to add a new book
to the database.
*/
func (storage DBStorage) InsertBook(book *Book) (int, error) {
	pg, err := sql.Open("postgres", getConnectionInfo())
	if err != nil {
		return -1, err
	}
	defer pg.Close()

	insertStatement := `
	INSERT INTO book (Title, Author, Publisher, Publish_Date, Rating, Status) 
	values ($1, $2, $3, $4, $5, $6)`

	_, err = pg.Exec(insertStatement, book.Title, book.Author, book.Publisher, book.PublishDate, book.Rating, book.Status)
	if err != nil {
		return -1, err
	}

	// Next, we retrieve the new entry's ID to provide the end user
	selectStatement := `
	SELECT ID 
	FROM book 
	WHERE TITLE = $1 
	AND AUTHOR = $2
	AND PUBLISHER = $3;`

	var id int

	err = pg.QueryRow(selectStatement, book.Title, book.Author, book.Publisher).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, err
}

/*
UpdateBook updates the entry with the given ID with the info in the given book JSON.
*/
func (storage DBStorage) UpdateBook(id int, book Book) error {
	pg, err := sql.Open("postgres", getConnectionInfo())
	if err != nil {
		return err
	}
	defer pg.Close()

	updateStatement := `
	update book
	set `
	updateCounter := 1

	params := []interface{}{id}

	// This is a hacky way around going back to learn/implement a full ORM system
	if book.IsValidTitle() {
		params = append(params, book.Title)
		updateStatement += buildUpdateStatement(&updateCounter, "title")
	}
	if book.IsValidAuthor() {
		params = append(params, book.Author)
		updateStatement += buildUpdateStatement(&updateCounter, "author")
	}
	if book.IsValidPublisher() {
		params = append(params, book.Publisher)
		updateStatement += buildUpdateStatement(&updateCounter, "publisher")
	}
	if book.IsValidPublishDate() {
		params = append(params, book.PublishDate)
		updateStatement += buildUpdateStatement(&updateCounter, "publish_date")
	}
	if book.IsValidRating() {
		params = append(params, fmt.Sprintf("%f", book.Rating))
		updateStatement += buildUpdateStatement(&updateCounter, "rating")
	}
	if book.Status != 0 {
		params = append(params, fmt.Sprintf("%d", book.Status))
		updateStatement += buildUpdateStatement(&updateCounter, "status")
	}

	if updateCounter == 0 {
		return nil
	}

	updateStatement += " where ID = $1;"

	log.Printf(`Update Statement: "%s"\n`, updateStatement)
	log.Printf(`Args: %v`, params)

	testExec(params)

	_, err = pg.Exec(updateStatement, params...)

	return err
}

func testExec(params ...interface{}) {
	log.Println("Got ", params)
}

func buildUpdateStatement(updates *int, key string) string {
	result := ""
	if *updates > 1 {
		result += ", "
	}
	*updates++
	result += fmt.Sprintf("%s = $%d", key, *updates)
	return result
}

/*
DeleteBook removes a book with given ID, if exists, from the database.
*/
func (storage DBStorage) DeleteBook(id int) error {
	pg, err := sql.Open("postgres", getConnectionInfo())
	if err != nil {
		return err
	}
	defer pg.Close()

	deleteStatement := `
	DELETE FROM book
	WHERE ID = $1;`

	_, err = pg.Exec(deleteStatement, id)
	if err != nil {
		return err
	}

	return err
}

/*
GetBook fetches the book entry with the provided ID number.
*/
func (storage DBStorage) GetBook(id int) (*Book, error) {
	var book Book

	pg, err := sql.Open("postgres", getConnectionInfo())
	if err != nil {
		return &book, err
	}
	defer pg.Close()

	selectStatement := `SELECT * FROM book WHERE ID = $1`
	err = pg.QueryRow(selectStatement, id).Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.PublishDate, &book.Rating, &book.Status)
	if err != nil {
		return &book, err
	}

	return &book, err
}

func getConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s "+
		"sslmode=disable",
		host, port, user, password, database)
}
