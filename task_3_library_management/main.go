package main

import (
	"bufio"
	"fmt"
	"os"
	"task_3_library_management/controllers"
	"task_3_library_management/services"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("----------------------------------------------------------------------------------------------------------")

		fmt.Println("\n------------------------- Library Management System -----------------------------")

		// fmt.Println("----------------------------------------------------------------")

		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("8. Remove Member")
		fmt.Println("9. Exit")

		fmt.Println("----------------------------------------------------------------------------------------------------------")

		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			controller.AddBook()
		case "2":
			controller.RemoveBook()
		case "3":
			controller.BorrowBook()
		case "4":
			controller.ReturnBook()
		case "5":
			controller.ListAvailableBooks()
		case "6":
			controller.ListBorrowedBooksByMember()
		case "7":
			controller.AddMember()
		case "8":
			controller.RemoveMember()
		case "9":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
