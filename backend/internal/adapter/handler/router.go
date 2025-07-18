package handler

import (
	"hr_management/internal/adapter/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	*fiber.App
}

type RouterParams struct {
	LeaveHandler *LeaveHandler
	Config      *config.HTTP
}

func NewRouter(p RouterParams) *Router {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     p.Config.AllowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	{
		leaves := api.Group("/leaves")
		{
			leaves.Post("/request", p.LeaveHandler.RequestLeave)
			leaves.Get("/all", p.LeaveHandler.GetAllLeaves)
			leaves.Put("/:id/approve", p.LeaveHandler.ApprovedLeave)
			leaves.Put("/:id/reject", p.LeaveHandler.RejectedLeave)
		}
	}
	return &Router{App: app}
}

func (r *Router) Start(port string) error {
	return r.Listen(port)
}
