package dbconnect

import (
	"database/sql"
	Config "api-booking-docter/config"
	_ "github.com/lib/pq" // Import driver PostgreSQL
	"log"
)

// Kết nối với cơ sở dữ liệu PostgreSQL
func ConnectDB() (*sql.DB, error) {
	cfg, err := Config.LoadConfig()
	// Cấu hình chuỗi kết nối PostgreSQL
	connStr := "user=" + cfg.PostSql.DBUser +
		" password=" + cfg.PostSql.DBPassword +
		" dbname=" + cfg.PostSql.DBName +
		" host=" + cfg.PostSql.DBHost +
		" port=" + cfg.PostSql.DBPort +
		" sslmode=disable" // SSL có thể được bật tùy theo cấu hình của bạn

	// Mở kết nối tới cơ sở dữ liệu
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database:", err)
		return nil, err
	}

	// In thông báo kết nối thành công
	log.Println("Successfully connected to the database!")
	return db, nil
}
