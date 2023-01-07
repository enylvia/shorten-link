package app

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

// Initialize PostgreSQL database

// func Connect() *gorm.DB {
// 	dsn := "host=localhost user=postgres password=1sampai10 dbname=belajar port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error connecting to database", err)
// 	}
// 	return db
// }

// Initialize Redis connection

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}
