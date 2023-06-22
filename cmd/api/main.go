package main

import (
	"github.com/farhodm/alif-test/internal/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("cannot load .env file:", err)
	}

	db, err := database.DBInit()
	if err != nil {
		log.Fatalln("cannot connect to DB:", err)
	}

	router := routes(db)
	if err := router.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalln("cannot run the server:", err)
	}
}
