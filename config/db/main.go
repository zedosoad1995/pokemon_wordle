package db_pkg

import (
	"fmt"
	"os"

	"github.com/zedosoad1995/pokemon-wordle/models/answer"
	"github.com/zedosoad1995/pokemon-wordle/models/board"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	"github.com/zedosoad1995/pokemon-wordle/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbUser, password, host, port, dbName, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&pokemon.Pokemon{})
	db.AutoMigrate(&board.Board{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&answer.Answer{})

	return db
}
