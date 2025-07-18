package handler

import (
	"hr_management/internal/adapter/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

type Router struct {
	*fiber.App
}

type RouterParams struct {
	LeaveHandler *LeaveHandler
	UserHandler  *UserHandler
	Config       *config.HTTP
}

func NewRouter(p RouterParams) *Router {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     p.Config.AllowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	// Swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// User routes
			users := v1.Group("/users")
			{
				users.Post("/register", p.UserHandler.RegisterEndpoint)
			}

			// Leave routes
			leaves := v1.Group("/leaves")
			{
				leaves.Post("/request", p.LeaveHandler.RequestLeave)
				leaves.Get("/all", p.LeaveHandler.GetAllLeaves)
				leaves.Put("/:id/approve", p.LeaveHandler.ApprovedLeave)
				leaves.Put("/:id/reject", p.LeaveHandler.RejectedLeave)
			}
		}

		user := v1.Group("/user")
		{
			user.Post("/register", p.UserHandler.RegisterEndpoint)
		}
	}
	return &Router{App: app}
}

func (r *Router) Start(port string) error {
	return r.Listen(port)
}
