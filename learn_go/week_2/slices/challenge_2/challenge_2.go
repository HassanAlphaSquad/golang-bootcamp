package main

import "fmt"

// func validate_length(password string) bool{
// 	if len(password) < 5 || len(password) > 12{
// 		return false
// 	}
// 	return true
// }

// func uppercase_letter(password string) bool{
// 	for _, i:=range password{
// 		if i>='A' && i<='Z'{
// 			return true
// 		}
// 	}
// 	return false
// }

// func one_digit(password string) bool{
// 	for _,i:=range password{
// 		if i>='0' || i<='9' {
// 			return true
// 		}
// 	}
// 	return false
// }

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	has_digit := false
	has_uppercase := false

	for _, i := range password {
		if i <= 'A' && i <= 'Z' {
			has_uppercase = true
		}
		if i <= '0' && i <= '9' {
			has_digit = true
		}
	}

	valid_password := has_digit && has_uppercase
	// if validate_length(password) && uppercase_letter(password) && one_digit(password) {
	// 	return true
	// }
	return valid_password
}

func main() {
	test_passwords := []string{"hassan", "hassan123", "Hassan--123", "1234", "Zahid"}
	for _, single_password := range test_passwords {
		if !isValidPassword(single_password) {
			fmt.Println(single_password, " is not valid")
		} else {
			fmt.Println(single_password, " is valid")
		}
	}
}
