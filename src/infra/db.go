package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	// ここでDBの接続設定を行う

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	fmt.Println(dsn)

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err !=nil {
		// ここでエラーが発生した場合はpanicを発生させる
		panic(err.Error())
		
		
	}

	return db


}
