package main

import (
	"hr_management/internal/adapter/config"
	"hr_management/internal/adapter/handler"
	"hr_management/internal/adapter/storage/mysql"
	"hr_management/internal/adapter/storage/mysql/repository"
	"hr_management/internal/core/service"
	"log"

	_ "hr_management/docs" // Swagger docs
)

// @title HRM Backend API
// @version 1.0
// @description Human Resource Management Backend API with Swagger documentation
// @host localhost:3000
// @BasePath /api/v1

func Init(config *config.Container) {
	db, err := mysql.NewDatabase(&config.DB)
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Initialize repositories
	leaveRepo := repository.NewGormLeaveRepo(db.DB)
	userRepo := repository.NewUserRepository(db.DB)

	// Initialize services
	leaveService := service.NewLeaveService(leaveRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	leaveHandler := handler.NewLeaveHandler(leaveService)
	userHandler := handler.NewUserHandler(userService)

	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := handler.NewRouter(handler.RouterParams{
		LeaveHandler: leaveHandler,
		UserHandler:  userHandler,
		Config:       &config.HTTP,
	})

	log.Println("Starting server on 127.0.0.1:3000")
	if err := router.Start("127.0.0.1:3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	Init(config)
}
