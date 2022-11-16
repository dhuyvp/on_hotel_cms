package device

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/app/queries"
	"hotel_cms/pkg/utils"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SearchDeviceData(db *sqlx.DB) fiber.Handler {
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
		HotelID := DataStruct.HotelID
		DeviceName := DataStruct.DeviceName
		IsActive := DataStruct.IsActive

		tableName := fmt.Sprintf("%s", os.Getenv("DEVICE"))
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}
		if DeviceName != nil {
			WHERECLause += " AND DeviceName LIKE '%" + *DeviceName + "%'"
		}

		var result []models.DeviceManage

		if HotelID == nil {
			log.Println("Don't have hotel that want to SEARCH device manage data")
			return c.Status(fiber.StatusOK).JSON(utils.Response{
				Success:    true,
				StatusCode: fiber.StatusOK,
				Data:       result,
			})
		}
		WHERECLause += " AND HotelID=?"
		queryDb += WHERECLause

		err := db.Select(&result, queryDb, *HotelID)

		if err != nil {
			log.Println("Error to SEARCH device manage data", err)
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
