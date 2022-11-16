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

func UpdateDeviceData(db *sqlx.DB) fiber.Handler {
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
		UpdateID := *DataStruct.DeviceID

		tableHotelName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryCheckHotel := "SELECT * FROM " + tableHotelName + " WHERE HotelID=?"
		var resultCheckHotel []models.HotelManage
		errCheck := db.Select(&resultCheckHotel, queryCheckHotel, *&DataStruct.HotelID)
		if errCheck != nil || resultCheckHotel == nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				Message:    "Don't have HotelID that I need update",
				StatusCode: fiber.StatusBadRequest,
			})
		}

		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryUpdate := utils.GetQueryUpdateDeviceManage(DataStruct)
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE DeviceID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE device manage data", errUpdate)
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
