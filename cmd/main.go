package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zedosoad1995/pokemon-wordle/config/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
}
