package controllers

import (
	"library_management/services"
)

type LibraryController struct {
	libraryService services.LibraryService
}

func NewLibraryController(service services.LibraryService) *LibraryController {
	return &LibraryController{
		libraryService: service,
	}
}

func (lc *LibraryController) AddBook(book services.Book) {

}
