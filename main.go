package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generatePassword(length int, includeNumbers bool, includeSymbols bool) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if includeNumbers {
		charset += "1234567890"
	}
	if includeSymbols {
		charset += "!@#$%^&*()_+-=[]\\{}|,./<>?"
	}

	//creates a random source everytime the program is opened
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[r.Intn(len(charset))] // r.Intn(len(charset)) is used to produce a random letter/number/or symbol
	} //												   and add it to the password slice

	return string(password)
}

func main() {
	var length int
	var includeNumbers string
	var includeSymbols string
	var addSymbols bool
	var addNums bool

	fmt.Println("How long do you want your password to be?")
	fmt.Scan(&length)

	fmt.Println("Do you want to include numbers? [Y/N]")
	fmt.Scan(&includeNumbers)

	includeNumbers = strings.ToLower(includeNumbers)
	if includeNumbers == "y" || includeNumbers == "yes" {
		addNums = true
	}

	fmt.Println("Do you want to include symbols? [Y/N]")
	fmt.Scan(&includeSymbols)

	includeSymbols = strings.ToLower(includeSymbols)
	if includeSymbols == "y" || includeSymbols == "yes" {
		addSymbols = true
	}

	password := generatePassword(length, addNums, addSymbols)

	var platform string
	fmt.Println("What platform is this password for?")
	fmt.Scan(&platform)

	entry := fmt.Sprintf("%s: %s\n", platform, password)

	fmt.Println()
	fmt.Printf("Your password for generated for %s is: %s\n", platform, password)

	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	_, err2 := file.WriteString(entry)
	if err2 != nil {
		fmt.Println("Error writing to file:", err2)
	}
}
