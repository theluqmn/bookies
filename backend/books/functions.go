

package books

import (
	"fmt"
	
	"main/util"
)

func addBook(id string, title string, author string, description string) bool {
	_, err := util.DB.Exec("INSERT INTO books (id, title, author, description) VALUES (?, ?, ?, ?)", id, title, author, description)
	if err != nil { fmt.Println(err); return false }
	
	return true
}

func bookExists(id string) bool {
	var count int
	err := util.DB.QueryRow("SELECT COUNT(*) FROM books WHERE id = ?", id).Scan(&count)
	if err != nil { return false }

	return count > 0
}