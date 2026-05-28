package users

import (
	"fmt"

	"main/util"
)

func signUp(id string, name string, password string) bool {
	fmt.Println(id, name, password)

	_, err := util.DB.Exec("INSERT INTO users (id, name, password) VALUES (?, ? ,?);", id, name, password)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
