package main

import "regexp"

/*
Book Basic book entry object
*/
type Book struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Publisher   string  `json:"publisher"`
	PublishDate string  `json:"publish_date"`
	Rating      float32 `json:"rating,omitempty"`
	Status      int     `json:"status,omitempty"`
}

/*
IsValidTitle returns true if the title is a non-empty string.
*/
func (book *Book) IsValidTitle() bool {
	return book.Title != ""
}

/*
IsValidAuthor returns true if the author is a non-empty string.
*/
func (book *Book) IsValidAuthor() bool {
	return book.Author != ""
}

/*
IsValidPublisher returns true if the publisher is a non-empty string.
*/
func (book *Book) IsValidPublisher() bool {
	return book.Publisher != ""
}

/*
IsValidPublishDate returns true if the publish date is in a MM-DD-YYYY format.
*/
func (book *Book) IsValidPublishDate() bool {
	re := regexp.MustCompile("^(0[1-9]|1[012])-(0[1-9]|[12][0-9]|[3][01])-([0-9]{4})$")
	return re.MatchString(book.PublishDate)
}

/*
IsValidRating returns true if the rating is a value between 1.0 and 3.0
*/
func (book *Book) IsValidRating() bool {
	return book.Rating >= 1.0 && book.Rating <= 3.0
}
