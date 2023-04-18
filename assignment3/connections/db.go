package connections

import (
	"github.com/joho/godotenv"
	"github.com/moldaspan/assignment3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	errr := godotenv.Load()

	if errr != nil {
		panic(errr)
	}

	dbURL := "host=db dbname=postgres user=postgres password=e!_sUltan747 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})
	DB = db
}
