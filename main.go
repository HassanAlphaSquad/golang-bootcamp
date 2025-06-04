package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome to Golang Bootcamp!")

	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	_, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	// dbQueries := database.New(db)
	age:=20
	if age>=18 {
		fmt.Printf("Yes, you are eligible for casting vote")
	}
}
