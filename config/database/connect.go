package database

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDb() {
	// dsn := "host=localhost user=postgres password='' dbname=go-db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                func() time.Time { return time.Now().Local() },
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		panic(err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = db
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	log.Println("running migrations")
	return DB.AutoMigrate(tables...)
}
