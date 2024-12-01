package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostSql PostSql
}

type PostSql struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	// Tìm và đọc file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		return nil, err
	}

	// Lấy giá trị từ file .env và gán vào cấu trúc Config
	config := &Config{
		PostSql: PostSql{
			DBHost:     os.Getenv("DB_HOST"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBPort:     os.Getenv("DB_PORT"),
		},
	}

	// Kiểm tra xem giá trị có tồn tại không (có thể bỏ qua nếu không cần thiết)
	if config.PostSql.DBHost == "" || config.PostSql.DBUser == "" || config.PostSql.DBPassword == "" || config.PostSql.DBName == "" || config.PostSql.DBPort == "" {
		log.Println("Warning: Some database environment variables are not set")
	}

	return config, nil
}
