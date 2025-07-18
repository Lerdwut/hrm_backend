package main

import (
	"hr_management/internal/adapter/config"
	mysql "hr_management/internal/adapter/storage/mysql"
	"log/slog"
	"os"
)

func Init(config *config.Container) {

	db, err := mysql.NewDatabase(&config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database migrated successfully")

	// Dependency injection
	//User

	// leaveRepo := leaveRepo.NewGormLeaveRepo()
	// leaveService := leaveService.NewLeaveService(leaveRepo)
	// leaveHandler := leaveHandler.NewLeaveHandler(leaveService)

	// app := fiber.New()
	// api := app.Group("/api")
	{
		// leaves := api.Group("/leaves")
		// {
		// 	leaves.Post("/request", leaveHandler.RequestLeave)
		// 	leaves.Get("/all", leaveHandler.GetAllLeaves)
		// 	leaves.Put("/:id/approve", leaveHandler.ApprovedLeave)
		// 	leaves.Put("/:id/reject", leaveHandler.RejectedLeave)
		// }
	}

	// log.Println("Starting server on :3000")
	// if err := app.Listen(":3000"); err != nil {
	// 	log.Fatal("Error starting server:", err)
	// }
}
