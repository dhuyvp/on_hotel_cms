package routes

import (
	"hotel_cms/app/controllers/hotel"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicHotelManageRoutes(db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/hotel_manage")

	route.Post("/insert", hotel.InsertHotelData(db))
	route.Get("/get", hotel.SearchHotelData(db))
	route.Put("/put", hotel.UpdateHotelData(db))
	route.Delete("/delete", hotel.DeleteHotelData(db))

}
