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

func InsertHotelData(db *sqlx.DB) fiber.Handler {
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

		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryColumns, queryValues := utils.GetColumnsAndValuesHotelManage(DataStruct)

		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT hotel manage data", errInsert)
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
