package routes

import (
	"hotel_cms/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicRoutes(app *fiber.App, db *sqlx.DB) {
	GroupAPI := app.Group("/api")

	GroupAPI.Post("/post", controllers.PostControl(db))

}
