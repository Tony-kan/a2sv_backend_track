package models

// Note : Status can be "available" or "borrowed"
type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}
