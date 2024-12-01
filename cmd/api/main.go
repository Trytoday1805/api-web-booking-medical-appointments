package main

import (
	"api-booking-docter/config"
	"api-booking-docter/internal/dbConnect"
	"fmt"
	"log"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// In config ra màn hình (debug)
	fmt.Printf("Load Config: %+v\n", cfg)

	// Sử dụng các giá trị trong cfg để kết nối database
	fmt.Printf("Connecting to DB: host=%s port=%s user=%s dbname=%s\n", cfg.PostSql.DBHost, cfg.PostSql.DBPort, cfg.PostSql.DBUser, cfg.PostSql.DBName)

	db, err := dbconnect.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()
}
