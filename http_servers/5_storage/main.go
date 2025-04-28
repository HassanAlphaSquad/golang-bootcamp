package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	_, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	// dbQueries := database.New(db)
	// defer db.Close()

	fmt.Println("Hello,  Storage")
}
