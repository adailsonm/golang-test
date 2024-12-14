package Database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Configuration struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildConfig() *Configuration {
	return &Configuration{
		Host:     os.Getenv("DB_HOST"),
		Port:     portConv(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func buildURL(configure *Configuration) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		configure.Host,
		configure.User,
		configure.Password,
		configure.DBName,
		configure.Port,
	)
}

func portConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Initialize() {
	config := buildConfig()
	url := buildURL(config)

	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	log.Println("connected")
}

func GetDB() *gorm.DB {
	return DB
}
