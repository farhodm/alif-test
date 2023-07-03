package main

import (
	"github.com/farhodm/ewallet/internal/database"
	"github.com/farhodm/ewallet/internal/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
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

	//creating default user
	// phone: +992111222333
	// password: qwerty
	user1 := models.User{
		Name:  "Vali",
		Phone: "+992111222333",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("qwerty"), 12)
	if err != nil {
		log.Fatalln("cannot generate password hash:", err)
	}
	user1.Password = string(hashedPassword)
	if err := db.Create(&user1).Error; err != nil {
		log.Fatalln("cannot create default user:", err)
	}
	//creating default user
	// phone: +992111444888
	// password: password
	user2 := models.User{
		Name:  "Ali",
		Phone: "+992111444888",
	}
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte("password"), 12)
	if err != nil {
		log.Fatalln("cannot generate password hash:", err)
	}
	user2.Password = string(hashedPassword)
	if err := db.Create(&user2).Error; err != nil {
		log.Fatalln("cannot create default user:", err)
	}

	wallet1 := models.Wallet{
		UserID: user1.ID,
	}
	if err := db.Create(&wallet1).Error; err != nil {
		log.Fatalln("cannot create wallet:", err)
	}
	wallet2 := models.Wallet{
		UserID: user2.ID,
		Type:   "identified",
	}
	if err := db.Create(&wallet2).Error; err != nil {
		log.Fatalln("cannot create wallet:", err)
	}

	log.Println("Successfully!")
}
