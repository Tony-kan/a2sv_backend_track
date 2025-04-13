package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"task_3_library_management/models"
	"task_3_library_management/services"
)

type LibraryController struct {
	libraryService services.LibraryService
}

func NewLibraryController(service services.LibraryService) *LibraryController {
	return &LibraryController{
		libraryService: service,
	}
}

func (libController *LibraryController) AddBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================================================")
	fmt.Print("Enter book Id : ")
	scanner.Scan()

	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter book Title : ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter book Author : ")
	scanner.Scan()
	author := scanner.Text()

	book := models.Book{
		ID:     bookId,
		Title:  title,
		Author: author,
		Status: "available",
	}

	err := libController.libraryService.AddBook(book)

	if err != nil {
		fmt.Println("\n=================================================================")
		fmt.Println("ERROR:", err)
		fmt.Println("=================================================================")
		return
	}

	fmt.Println("=================================================================")

	fmt.Println("Book added successfully!")

	fmt.Println("=================================================================")

}

func (libController *LibraryController) ListAvailableBooks() {
	books := libController.libraryService.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No boooks available")
		return
	}
	fmt.Println("Books available : ")
	for _, book := range books {
		fmt.Printf("Id : %d | Title : %s | Author : %s | Status : %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

func (libController *LibraryController) RemoveBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================================================")
	fmt.Print("Enter book Id of the book to remove : ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	libController.libraryService.RemoveBook(bookId)
	fmt.Println("=================================================================")

	fmt.Printf("Book with Id : %d has been removed successfully!", bookId)

	fmt.Println("=================================================================")

}

func (libController *LibraryController) BorrowBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================================================")
	fmt.Print("Enter book Id of the book to borrow : ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member Id : ")
	scanner.Scan()
	memberId, _ := strconv.Atoi(scanner.Text())

	err := libController.libraryService.BorrowBook(bookId, memberId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("=================================================================")

	fmt.Printf("Book with Id : %d has been borrowed successfully!", bookId)

	fmt.Println("=================================================================")

}

func (libController *LibraryController) ReturnBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================================================")
	fmt.Print("Enter book Id of the book to return : ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member Id : ")
	scanner.Scan()
	memberId, _ := strconv.Atoi(scanner.Text())

	err := libController.libraryService.ReturnBook(bookId, memberId)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("=================================================================")

	fmt.Printf("Book with Id : %d has been returned successfully!", bookId)

	fmt.Println("=================================================================")

}

func (libController *LibraryController) ListBorrowedBooksByMember() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=================================================================")

	fmt.Print("Enter member Id : ")
	scanner.Scan()
	memberId, _ := strconv.Atoi(scanner.Text())

	books := libController.libraryService.ListBorrowedBooks(memberId)

	if len(books) == 0 {
		fmt.Println("No books borrowed by member")
		return
	}
	fmt.Printf("Books borrowed by member with Id %d : \n", memberId)
	for _, book := range books {
		fmt.Printf("Id : %d | Title : %s | Author : %s | Status : %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

func (libController *LibraryController) AddMember() {}

func (libController *LibraryController) RemoveMember() {}
