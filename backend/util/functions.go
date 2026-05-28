// reusable utility functions

package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) string {
	inputBytes, _ := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	return string(inputBytes)
}

func InputLongEnough(input string, min int, max int) bool {
	length := len(input)
	return length >= min && length <= max
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GenerateRandomID(size int) string {
	b := make([]byte, size)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}