package main

import (
	"hr_management/internal/adapter/config"
	"hr_management/internal/adapter/handler"
	mongodb "hr_management/internal/adapter/storage/mongodb"
	"hr_management/internal/adapter/storage/mongodb/repository"
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
	// Connect to MongoDB
	db, err := mongodb.NewDatabase(&config.DB)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer db.Close()

	leaveRepo := repository.NewMongoLeaveRepo(db.DB)
	leaveService := service.NewLeaveService(leaveRepo)
	leaveHandler := handler.NewLeaveHandler(leaveService)

	// userRepo := repository.NewUserRepository(db.DB)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	router := handler.NewRouter(handler.RouterParams{
		LeaveHandler: leaveHandler,
		// UserHandler:  userHandler,
		Config:       &config.HTTP,
	})

	log.Println("Starting server on 127.0.0.1:3000")
	if err := router.Start("127.0.0.1:3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	Init(cfg)
}
