package hotel

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

func SearchHotelData(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var HotelName *string
		var IsActive *bool
		p := queries.GetModelHotel(c)

		if p != nil {
			DataStruct := *p
			HotelName = DataStruct.HotelName
			IsActive = DataStruct.IsActive
		}
		tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"

		if HotelName != nil {
			WHERECLause += " AND HotelName LIKE '%" + *HotelName + "%'"
		}
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}

		queryDb += WHERECLause
		var result []models.HotelManage
		err := db.Select(&result, queryDb)

		if err != nil {
			log.Println("Error to SEARCH hotel manage data", err)
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
