package users

import "fmt"

func signUp(id string, name string, password string) bool {
	fmt.Println(id, name, password)
	return true
}
