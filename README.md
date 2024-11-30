# api-web-booking-medical-appointments
# api-web-booking-medical-appointments
# Clean Architecture
backend/
├── cmd/                     # Entry points của ứng dụng (main.go)
│   └── api/
│       └── main.go          # Điểm bắt đầu chạy ứng dụng
├── config/                  # Cấu hình ứng dụng (file, env, v.v.)
│   ├── config.go            # Cấu hình chính (load file .env, yaml)
│   ├── app.yaml             # File config app (nếu cần)
│   └── db.yaml              # File config database (nếu cần)
├── internal/                # Thư mục chính của ứng dụng
│   ├── domain/              # Entities và Business Logic (Domain Layer)
│   │   ├── user.go          # Entity User
│   │   ├── patient.go       # Entity Patient
│   │   └── appointment.go   # Entity Appointment
│   ├── usecase/             # Application Logic (Usecase Layer)
│   │   ├── user_usecase.go  # Usecase liên quan đến User
│   │   ├── patient_usecase.go
│   │   ├── appointment_usecase.go
│   │   └── interfaces.go    # Interfaces định nghĩa giao tiếp với repository
│   ├── repository/          # Data Access Layer (Repository Layer)
│   │   ├── user_repo.go     # Repository User
│   │   ├── patient_repo.go  # Repository Patient
│   │   ├── appointment_repo.go
│   │   └── mock/            # Mock Repository (cho test)
│   │       ├── user_repo_mock.go
│   │       └── patient_repo_mock.go
│   ├── handler/             # Giao tiếp với framework, xử lý request/response
│   │   ├── user_handler.go  # Handler cho User
│   │   ├── patient_handler.go
│   │   ├── appointment_handler.go
│   │   └── middleware/      # Middleware cho request (nếu có)
│   │       ├── auth.go      # Middleware xác thực
│   │       └── logger.go    # Middleware logging
│   ├── service/             # Các service bên ngoài (gửi email, push notification, v.v.)
│   │   ├── email_service.go # Service gửi email
│   │   ├── notification_service.go
│   │   └── logger_service.go
│   ├── dto/                 # Data Transfer Objects (giữa layers)
│   │   ├── user_dto.go
│   │   ├── patient_dto.go
│   │   └── appointment_dto.go
│   └── error/               # Xử lý lỗi
│       ├── error.go         # Định nghĩa lỗi chung
│       └── custom_error.go  # Định nghĩa lỗi đặc biệt
├── pkg/                     # Các thư viện hoặc module có thể tái sử dụng
│   ├── utils/               # Utility functions (helpers)
│   │   ├── validation.go    # Hàm validate dữ liệu
│   │   ├── date.go          # Hàm xử lý thời gian
│   │   └── hash.go          # Hàm hash dữ liệu (mật khẩu, v.v.)
│   ├── logger/              # Logger cho ứng dụng
│   │   ├── logger.go
│   │   └── formatter.go
│   ├── database/            # Kết nối với database
│   │   ├── postgres.go      # Kết nối PostgreSQL
│   │   ├── mysql.go         # Kết nối MySQL
│   │   └── mongo.go         # Kết nối MongoDB
│   ├── jwt/                 # Xử lý JWT (xác thực)
│   │   ├── jwt.go
│   │   └── jwt_utils.go
│   └── cache/               # Cache (Redis, v.v.)
│       ├── redis.go
│       └── memory_cache.go
├── test/                    # Thư mục test (integration, e2e)
│   ├── integration/         # Test tích hợp
│   │   ├── user_integration_test.go
│   │   └── appointment_integration_test.go
│   └── e2e/                 # Test end-to-end
│       └── e2e_test.go
├── go.mod                   # File quản lý dependencies (module Go)
├── go.sum                   # File lock dependencies
└── README.md                # Hướng dẫn và thông tin dự án
