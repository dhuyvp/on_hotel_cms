package hotel

import (
	"fmt"
	"hotel_cms/app/queries"
	"hotel_cms/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UpdateHotelData(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := queries.GetModelHotel(c)

		if p == nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				Message:    "No data",
				StatusCode: fiber.StatusBadRequest,
			})
		}

		DataStruct := *p
		UpdateID := DataStruct.HotelID

		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryUpdate := utils.GetQueryUpdateHotelManage(DataStruct)

		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE HotelID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE hotel manage data", errUpdate)
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
