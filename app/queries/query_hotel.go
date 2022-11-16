package queries

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func GetModelHotel(c *fiber.Ctx) *models.HotelManage {
	var p *models.HotelManage

	if err := c.BodyParser(p); err != nil {
		fmt.Println(err)
		return nil
	}

	*p = middleware.CheckHotelInformation(*p)

	return p
}
