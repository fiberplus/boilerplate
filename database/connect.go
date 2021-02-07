package database

import (
	model "boilerplate/models"
	"fmt"
	"os"
	"sync"

	config "boilerplate/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

// DriverMySQL - MySQL connection grouping
type DriverMySQL struct {
	conn *gorm.DB
	// more connection here
}

var db *DriverMySQL

// GetDB - Create Db singleton once, return same singleton db after, default connection,
// copy this for another conn
func GetDB() *gorm.DB {
	once.Do(func() {
		fmt.Println("Mysql singleton creation, This message should appear once")
		db = &DriverMySQL{conn: ConnectDB()}
	})
	return db.conn
}

// ConnectDB connect to db
func ConnectDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.App,
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	if err != nil {
		panic(err)
	}

	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")

	dbconn.AutoMigrate(&model.User{})

	fmt.Println("Database Migrated")
	return dbconn
}
