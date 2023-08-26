package database

import (
	"os"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	host	 := os.Getenv("DB_HOST")
	port	 := os.Getenv("DB_PORT")
	user	 := os.Getenv("DB_USER")
	dbname	 := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")

	DB = Dbinstance{
		Db: db,
	}
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	return DB.Db.AutoMigrate(tables...)
}
