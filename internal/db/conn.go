package db

import (
	"fmt"
	"log"

	"github.com/captainGeech42/hotline/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbHandle *gorm.DB = nil

// https://gorm.io/docs/connecting_to_the_database.html
func ConnectToDb(cfg config.Database) bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("successfully connected to database")

	db.AutoMigrate(&Callback{}, &HttpRequest{}, &DnsRequest{})
	log.Println("ran db migrations")

	dbHandle = db
	return true
}
