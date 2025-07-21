package main

import (
	"hr_management/internal/adapter/config"
	"hr_management/internal/adapter/handler"
	"hr_management/internal/adapter/storage/mongo"
	"hr_management/internal/adapter/storage/mongo/repository"
	"hr_management/internal/core/service"
	"log"

	_ "hr_management/docs" // Swagger docs
)

// @title HRM Backend API
// @version 1.0
// @description Human Resource Management Backend API with Swagger documentation
// @host localhost:3000
// @BasePath /api/v1

// func Init(config *config.Container) {
// 	db, err := mysql.NewDatabase(&config.DB)
// 	if err != nil {
// 		log.Fatalf("Error initializing database connection: %v", err)
// 	}

// 	err = db.Migrate()
// 	if err != nil {
// 		log.Fatalf("Error migrating database: %v", err)
// 	}

// 	leaveRepo := repository.NewGormLeaveRepo(db.DB)
// 	leaveService := service.NewLeaveService(leaveRepo)
// 	leaveHandler := handler.NewLeaveHandler(leaveService)

// 	userRepo := repository.NewUserRepository(db.DB)
// 	userService := service.NewUserService(userRepo)
// 	userHandler := handler.NewUserHandler(userService)

// 	router := handler.NewRouter(handler.RouterParams{
// 		LeaveHandler: leaveHandler,
// 		UserHandler:  userHandler,
// 		Config:       &config.HTTP,
// 	})

// 	log.Println("Starting server on 127.0.0.1:3000")
// 	if err := router.Start("127.0.0.1:3000"); err != nil {
// 		log.Fatal("Error starting server:", err)
// 	}
// }

func Init(config *config.Container) {
	db, err := mongo.ConnectMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer db.Close()

	// Create dependencies
	userRepo := repository.NewMongoUserRepository(db.Database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := handler.NewRouter(handler.RouterParams{
		UserHandler: userHandler,
		Config:      &config.HTTP,
	})

	log.Printf("Starting server on :%s", config.HTTP.Port)
	if err := router.Start(config.HTTP.URL + ":" + config.HTTP.Port); err != nil {
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
