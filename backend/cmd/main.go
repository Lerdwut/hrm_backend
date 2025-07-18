package main

import (
	"hr_management/internal/adapter/config"
	"hr_management/internal/adapter/handler"
	"hr_management/internal/adapter/storage/mysql"
	"hr_management/internal/adapter/storage/mysql/repository"
	"hr_management/internal/core/service"
	"log"
)

func Init(config *config.Container) {
	db, err := mysql.NewDatabase(&config.DB)
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	leaveRepo := repository.NewGormLeaveRepo(db.DB)
	leaveService := service.NewLeaveService(leaveRepo)
	leaveHandler := handler.NewLeaveHandler(leaveService)

	router := handler.NewRouter(handler.RouterParams{
		LeaveHandler: leaveHandler,
	})

	log.Println("Starting server on :3000")
	if err := router.Start(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func main() {

}
