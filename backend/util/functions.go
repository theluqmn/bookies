package util

import (
	"os"
	"os/exec"

	"golang.org/x/crypto/bcrypt"
)

// hashing function
func Hash(input string) string {
	inputBytes, _ := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	return string(inputBytes)
}

// checks if input is within specified min/max length
func InputLongEnough(input string, min int, max int) bool {
	length := len(input)
	return length >= min && length <= max
}

// clears the terminal
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
