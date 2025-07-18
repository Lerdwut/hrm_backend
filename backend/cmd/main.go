package main

import (
	"github.com/gofiber/fiber/v2"
	leaveHandler "hr_management/internal/adapter/handler"
	leaveRepo "hr_management/internal/adapter/storage/mysql/repository"
	leaveService "hr_management/internal/core/service"
	"log"
)

func main() {
	leaveRepo := leaveRepo.NewGormLeaveRepo()
	leaveService := leaveService.NewLeaveService(leaveRepo)
	leaveHandler := leaveHandler.NewLeaveHandler(leaveService)

	app := fiber.New()
	api := app.Group("/api")
	{
		leaves := api.Group("/leaves")
		{
			leaves.Post("/request", leaveHandler.RequestLeave)
			leaves.Get("/all", leaveHandler.GetAllLeaves)
			leaves.Put("/:id/approve", leaveHandler.ApprovedLeave)
			leaves.Put("/:id/reject", leaveHandler.RejectedLeave)
		}
	}

	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
