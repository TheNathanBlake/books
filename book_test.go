package main

import "testing"

var titleTest = []struct {
	title   string
	isValid bool
}{
	{"A Big Long String Or Something", true},
	{"Onie", true},
	{"A", true},
	{"", false},
}

func TestTitleValidity(t *testing.T) {
	var book Book
	for _, test := range titleTest {
		book = Book{Title: test.title}
		if book.IsValidTitle() != test.isValid {
			t.Errorf(`expected "%s" isValid to be %v`, book.Title, test.isValid)
		}
	}
}

var authorTest = []struct {
	author  string
	isValid bool
}{
	{"A Big Long Name Or Something", true},
	{"Onie", true},
	{"A", true},
	{"", false},
}

func TestAuthorValidity(t *testing.T) {
	var book Book
	for _, test := range authorTest {
		book = Book{Author: test.author}
		if book.IsValidAuthor() != test.isValid {
			t.Errorf(`expected "%s" isValid to be %v`, book.Author, test.isValid)
		}
	}
}

var publisherTest = []struct {
	publisher string
	isValid   bool
}{
	{"A Big Long Name Or Something", true},
	{"Onie", true},
	{"A", true},
	{"", false},
}

func TestPublisherValidity(t *testing.T) {
	var book Book
	for _, test := range publisherTest {
		book = Book{Publisher: test.publisher}
		if book.IsValidPublisher() != test.isValid {
			t.Errorf(`expected "%s" isValid to be %v`, book.Publisher, test.isValid)
		}
	}
}

var publishDateTest = []struct {
	publishDate string
	isValid     bool
}{
	{"01-01-2018", true},
	{"12-31-2000", true},
	{"11-11-9000", true},
	{"13-12-2000", false},
	{"1-01-2019", false},
	{"0-12-2000", false},
	{"11-00-2000", false},
	{"11-33-2000", false},
	{"11-333-2000", false},
	{"11-11-10000", false},
	{"11-11-20X6", false},
}

func TestPublishDateValidity(t *testing.T) {
	var book Book
	for _, test := range publishDateTest {
		book = Book{PublishDate: test.publishDate}
		if book.IsValidPublishDate() != test.isValid {
			t.Errorf(`expected "%s" isValid to be %v`, book.PublishDate, test.isValid)
		}
	}
}

func TestRatingValidity(t *testing.T) {
	var validRatings = []float32{1.0, 1.11111, 2.0, 2.22222, 2.99999, 3.0}
	var invalidRatings = []float32{0.99999, 0.0, 3.00001, -1.1}
	var book Book
	for _, rating := range validRatings {
		book = Book{Rating: rating}
		if !book.IsValidRating() {
			t.Errorf("expected %f to be a valid rating", rating)
		}
	}
	for _, rating := range invalidRatings {
		book = Book{Rating: rating}
		if book.IsValidRating() {
			t.Errorf("expected %f to be an invalid rating", rating)
		}
	}

}
