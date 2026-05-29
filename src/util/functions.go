// reusable utility functions

package util

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"os/exec"
	"time"

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
		LogError(err)
	}
	
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

func GetFormattedTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.00")
}