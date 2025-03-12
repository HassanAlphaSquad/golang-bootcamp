package main

import (
	"fmt"
)

func validate_length(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	return true
}

func uppercase_letter(password string) bool {
	for _, i := range password {
		if i >= 'A' && i <= 'Z' {
			return true
		}
	}
	return false
}

func one_digit(password string) bool {
	for _, i := range password {
		if i >= 0 || i <= 9 {
			return true
		}
	}
	return false
}

func isValidPassword(password string) bool {
	if validate_length(password) && uppercase_letter(password) && one_digit(password) {
		return true
	}
	return false
}

func main() {
	if isValidPassword("Hassan") {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}
}
