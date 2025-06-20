package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English)

func getName(fName, lName string) string {
	return titleCaser.String(fName) + " " + titleCaser.String(lName)
}

func eligibility(age int) bool {
	return age >= 18 && age <= 120
}

func ternaryCondition(age int, eligibility func(int) bool, trueValue, falseValue string) string {
	if eligibility(age) {
		return trueValue
	}
	return falseValue
}

func profile(getName func(string, string) string, eligibility func(int) bool, first, last string, age int) string {
	name := getName(first, last)
	status := ternaryCondition(age, eligibility, "Eligible", "Ineligible")
	return fmt.Sprintf("%v, you are %v", name, status)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("*********************************")
	fmt.Println("Welcome to the Eligibility Checker")
	fmt.Println("Type 'q' at any point to quit.")
	fmt.Println("*********************************")

	for {
		fmt.Print("Enter First Name: ")
		firstName, _ := reader.ReadString('\n')
		firstName = strings.TrimSpace(firstName)
		if strings.ToLower(firstName) == "q" {
			fmt.Println("Exiting...")
			break
		}

		fmt.Print("Enter Last Name: ")
		lastName, _ := reader.ReadString('\n')
		lastName = strings.TrimSpace(lastName)
		if strings.ToLower(lastName) == "q" {
			fmt.Println("Exiting...")
			break
		}

		var age int
		for {
			fmt.Print("Enter Age: ")
			ageInput, _ := reader.ReadString('\n')
			ageInput = strings.TrimSpace(ageInput)
			if strings.ToLower(ageInput) == "q" {
				fmt.Println("Exiting...")
				return
			}
			parsedAge, err := strconv.Atoi(ageInput)
			if err != nil || parsedAge <= 0 {
				fmt.Println("Invalid age. Please enter a valid positive number.")
				continue
			}
			age = parsedAge
			break
		}

		fmt.Println("---------------------------------")
		fmt.Println(profile(getName, eligibility, firstName, lastName, age))
		fmt.Println("---------------------------------")
	}
}
