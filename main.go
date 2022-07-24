package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+=~`.")

func randSeq(index int) string {
	b := make([]rune, index)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	var passwordLength int
	fmt.Print("Hello, Please enter password length you want: ")
	fmt.Scanln(&passwordLength)
	rand.Seed(time.Now().UnixNano()) //seed it unique

	fmt.Println("Generated password:", randSeq(passwordLength))
	fmt.Println("Press any key to close")
	fmt.Scanln()
}
