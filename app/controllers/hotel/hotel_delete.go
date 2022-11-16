package hotel

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

func DeleteHotelData(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var deleteID int

		deleteID, errConvert := strconv.Atoi(c.FormValue("deleteID"))
		if errConvert != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				Message:    "Error when convert string to int ",
				StatusCode: fiber.StatusBadRequest,
			})
		}

		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryDb := "DELETE FROM " + tableName + " WHERE HotelID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE HotelID=?"

		var result []models.HotelManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT hotel manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE hotel manage data", errDelete)
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
