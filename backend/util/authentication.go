// functions for user authentication

package util

import (
	"fmt"
	"crypto/rand"
	"encoding/base64"
)

func SessionTokenCreate(id string) string {
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.StdEncoding.EncodeToString(b)

	_, err := DB.Exec("INSERT INTO sessions (id, token) VALUES (?, ?);", id, token)
	if err != nil { fmt.Println(err); return "" }
	
	return token
}

func SessionTokenVerify(token string) string {
	var id string
	err := DB.QueryRow("SELECT id FROM sessions WHERE token = ?;", token).Scan(&id)
	if err != nil { return "" }
	return id
}
