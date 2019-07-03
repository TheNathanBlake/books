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

func (book *Book) IsValidTitle() bool {
	return book.Title != ""
}

func (book *Book) IsValidAuthor() bool {
	return book.Author != ""
}

func (book *Book) IsValidPublisher() bool {
	return book.Publisher != ""
}

func (book *Book) IsValidPublishDate() bool {
	re := regexp.MustCompile("^(0[1-9]|1[012])-(0[1-9]|[12][0-9]|[3][01])-([0-9]{4})$")
	return re.MatchString(book.PublishDate)
}

func (book *Book) IsValidRating() bool {
	return book.Rating >= 1.0 && book.Rating <= 3.0
}
