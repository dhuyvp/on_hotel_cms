package routes

import (
	"hotel_cms/app/controllers/device"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicDeviceManageRoutes(db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/device_manage")

	route.Post("/insert", device.InsertDeviceData(db))
	route.Get("/get", device.SearchDeviceData(db))
	route.Put("/put", device.UpdateDeviceData(db))
	route.Delete("/delete", device.DeleteDeviceData(db))

}
