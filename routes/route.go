package routes

import (
	"amquizdua/controllers"
	"amquizdua/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	user := r.Group("/api")

	user.Get("/", middlewares.Auth, controllers.Index)
	user.Get("/:id", middlewares.Auth, controllers.Show)
	user.Post("/", middlewares.Auth, controllers.Create)
	user.Put("/:id", middlewares.Auth, controllers.Update)
}