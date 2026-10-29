// โค้ด คำสั่้ง Connect Database

package main

import (
	"fmt"
	"log"
	"os"
	"time"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host = "localhost"
	port = 3306
	username = "root"
	password = ""
	dbname = "sau_db"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200 * time.Microsecond,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Printf("DB connect successful")

	if err := db.AutoMigrate(&Student{}); err != nil{
		panic("failed to migrated db")
	}

	fmt.Printf("Migration success")
	return db, nil
}