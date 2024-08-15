package database

import (
	"log"

	"github.com/matiaseiglesias/storiChallenge/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	Db *gorm.DB
}

func Innit(config *config.Database) *DataBase {
	user := config.UserID
	pass := config.Pass
	host := config.Host
	port := config.Port
	schema := config.Schema

	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + schema + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error: ", err)
		panic("failed to connect database")
	}

	// err = database.AutoMigrate(&models.NFT{}, &models.User{}, &models.Image{})

	if err != nil {
		log.Println("Error: ", err)
		panic("failed to migrate database")
	}
	return &DataBase{
		Db: database,
	}
}
