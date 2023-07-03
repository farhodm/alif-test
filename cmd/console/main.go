package main

import (
	"github.com/farhodm/ewallet/internal/database"
	"github.com/farhodm/ewallet/internal/models"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("cannot load .env file:", err)
	}

	db, err := database.DBInit()
	if err != nil {
		log.Fatalln("cannot connect to DB:", err)
	}
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Wallet{})
	db.Migrator().DropTable(&models.Transaction{})

	if err := db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	); err != nil {
		log.Fatalln("cannot re-migrate the DB:", err)
	}
	for i := 0; i <= 10; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)
		if err != nil {
			log.Fatalln("cannot generate password hash:", err)
		}

		user := models.User{
			Name:     faker.FirstName(),
			Phone:    faker.Phonenumber(),
			Password: string(hashedPassword),
			Wallet:   &models.Wallet{},
		}

		if rand.Int()%2 == 0 {
			user.Wallet.Type = "identified"
		} else {
			user.Wallet.Type = "non-identified"
		}

		if user.Wallet.Type == "identified" {
			user.Wallet.Balance = uint64(rand.Intn(100_000_00))
		} else {
			user.Wallet.Balance = uint64(rand.Intn(10_000_00))
		}

		if err := db.Create(&user).Error; err != nil {
			log.Fatalln("cannot create user:", err)
		}
	}
	log.Println("Successfully!")
}
