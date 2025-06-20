package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divisor can't be zero")
	}
	return x / y, nil
}

func maths(x, y int, op string, resultChan chan string) {
	var result int
	switch op {
	case "+":
		result = add(x, y)
	case "-":
		result = sub(x, y)
	case "*":
		result = mul(x, y)
	case "/":
		divResult, err := div(x, y)
		if err != nil {
			resultChan <- fmt.Sprintf("Error: %v", err)
			return
		}
		result = divResult
	default:
		resultChan <- "Invalid operation"
		return
	}
	resultChan <- fmt.Sprintf("%d %s %d = %d", x, op, y, result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	resultChan := make(chan string)

	fmt.Println("Welcome to the concurrent calculator!")
	fmt.Println("Type 'q' at any prompt to quit.")
	fmt.Println("-------------------------------------")

	for {
		var a, b int
		var op string

		fmt.Print("Enter first value: ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "q" {
			fmt.Println("Exiting...")
			break
		}
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		a = val

		fmt.Print("Enter second value: ")
		if !scanner.Scan() {
			break
		}
		input = strings.TrimSpace(scanner.Text())
		if input == "q" {
			fmt.Println("Exiting...")
			break
		}
		val, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		b = val

		fmt.Print("Enter operation (+, -, *, /): ")
		if !scanner.Scan() {
			break
		}
		op = strings.TrimSpace(scanner.Text())
		if op == "q" {
			fmt.Println("Exiting...")
			break
		}

		go maths(a, b, op, resultChan)

		fmt.Println(<-resultChan)
		fmt.Println("-------------------------------------")
	}
}
