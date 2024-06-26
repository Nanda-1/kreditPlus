package database

import (
	"kreditPlus/app/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// OpenDb opens a connection to the MySQL database using GORM.
func OpenDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Customer{}, &models.Tenor{}, &models.Limits{}, &models.Loan{}, &models.Notification{})
	DB = db
	return DB, nil
}

func ConnectToDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/kredit_plus?charset=utf8mb4&parseTime=True&loc=Local"
	var counts int

	for {
		connection, err := OpenDb(dsn)
		if err != nil {
			log.Println("MySQL not yet ready ...")
			counts++
		} else {
			log.Println("Connected to MySQL!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
	}
}
