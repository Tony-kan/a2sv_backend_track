package services

import (
	"fmt"
	"task_3_library_management/models"
)

type LibraryService interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) error {
	if _, exists := l.books[book.ID]; exists {
		return fmt.Errorf("book with ID %d already exists", book.ID)
	}

	l.books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.books[bookID]; !exists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	if book, exists := l.books[bookID]; exists {
		if book.Status == "borrowed" {
			return fmt.Errorf("book with ID %d is already borrowed", bookID)
		}
		book.Status = "borrowed"
		l.books[bookID] = book
		l.members[memberID].BorrowedBooks = append(l.members[memberID].BorrowedBooks, book)
		return nil
	}
	return fmt.Errorf("book with ID %d does not exist", bookID)
}
func (l *Library) ReturnBook(bookID int, memberID int) error {
	if book, exists := l.books[bookID]; exists {
		if book.Status == "available" {
			return fmt.Errorf("book with ID %d is already available", bookID)
		}
		book.Status = "available"
		l.books[bookID] = book
		for i, borrowedBook := range l.members[memberID].BorrowedBooks {
			if borrowedBook.ID == bookID {
				l.members[memberID].BorrowedBooks = append(l.members[memberID].BorrowedBooks[:i], l.members[memberID].BorrowedBooks[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("book with ID %d is not borrowed by member with ID %d", bookID, memberID)
	}
	return fmt.Errorf("book with ID %d does not exist", bookID)
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	if member, exists := l.members[memberID]; exists {
		return member.BorrowedBooks
	}
	return nil
}
