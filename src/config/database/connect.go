package database

import (
	"gofiber-todo/src/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDb() {

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// 	NowFunc:                func() time.Time { return time.Now().Local() },
	// 	Logger:                 logger.Default.LogMode(logger.Info),
	// })

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB + " port=" + config.DB_PORT + " sslmode=disable TimeZone=Asia/Seoul",
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                func() time.Time { return time.Now().Local() },
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		panic(err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Error)

	DB = db
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	log.Println("running migrations")
	return DB.AutoMigrate(tables...)
}
