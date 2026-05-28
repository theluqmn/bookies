// user-related functions mostly relating to SQL

package users

import (
	"fmt"

	"main/util"

	"golang.org/x/crypto/bcrypt"
)

func signUp(id string, name string, password string) bool {
	fmt.Println(id, name, password)

	_, err := util.DB.Exec("INSERT INTO users (id, name, password) VALUES (?, ? ,?);", id, name, password)
	if err != nil { fmt.Println(err); return false }

	return true
}

func userExists(id string) bool {
	var count int
	err := util.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", id).Scan(&count)
	if err != nil { return false }

	return count > 0
}

func comparePassword(id string, password string) bool {
	var hashed string
	err := util.DB.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&hashed)
	if err != nil { return false }

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil { return false }

	return true
}
