// reusable utility functions

package util

import (
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