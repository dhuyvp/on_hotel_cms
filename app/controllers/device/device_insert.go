package device

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/app/queries"
	"hotel_cms/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func InsertDeviceData(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := queries.GetModelDevice(c)

		if p == nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				Message:    "No data",
				StatusCode: fiber.StatusBadRequest,
			})
		}

		DataStruct := *p
		HotelID := *(DataStruct.HotelID)

		tableName := fmt.Sprintf("%s", os.Getenv("DEVICE"))
		tableHotelName := fmt.Sprintf("%s", os.Getenv("HOTEL"))

		queryColumns, queryValues := utils.GetColumnsAndValuesDeviceManage(DataStruct)
		var result []models.HotelManage
		queryHotel := "SELECT * FROM " + tableHotelName + " WHERE HotelID=? LIMIT 1"

		err := db.Select(&result, queryHotel, HotelID)
		if err != nil || result == nil {
			log.Println("Error to SELECT Hotel that want to INSERT device manage data otel with HotelID = ", HotelID, err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		if *result[0].DevicesNumber >= *result[0].DevicesLimit {
			log.Println("Enough devices", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT hotel manage data", errInsert)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		queryUpdateHotel := "UPDATE " + tableHotelName + " SET DevicesNumber=? WHERE HotelID=?"
		_, errUpdateHotel := db.Exec(queryUpdateHotel, *result[0].DevicesNumber+1, HotelID)
		if errUpdateHotel != nil {
			log.Println("Update Hotel data is denied", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       DataStruct,
		})
	}
}
