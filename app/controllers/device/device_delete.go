package device

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/utils"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func DeleteDeviceData(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		deleteID, errConvert := strconv.Atoi(c.FormValue("DeviceID"))
		if errConvert != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				Message:    "Error when convert string to int ",
				StatusCode: fiber.StatusBadRequest,
			})
		}

		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryDb := "DELETE FROM " + tableName + " WHERE DeviceID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE DeviceID=?"

		var result []models.DeviceManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT device manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE device manage data", errDelete)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		var resultHotel []models.HotelManage
		tableHotelName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryUpdateHotel := "UPDATE " + tableHotelName + " SET DevicesNumber=? WHERE HotelID=?"
		querySelectHotel := "SELECT * FROM " + tableHotelName + " WHERE HotelID=?"
		errSelectHotel := db.Select(&result, querySelectHotel, *result[0].HotelID)
		if errSelectHotel != nil {
			log.Println("Error SELECT hotel manage data that needs to be UPDATE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errUpdateHotel := db.Exec(queryUpdateHotel, *resultHotel[0].DevicesNumber-1, *result[0].HotelID)
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
			Data:       result,
		})
	}
}
