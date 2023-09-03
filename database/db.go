package database

import (
	"log"

	"github.com/joserafaelSH/go-gorm-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func DataBaseConnect() {
	connectString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("Could not connect to the database")
	}

	Db.AutoMigrate(&models.Aluno{})
}