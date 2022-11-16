package queries

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func GetModelDevice(c *fiber.Ctx) *models.DeviceManage {
	var p *models.DeviceManage

	if err := c.BodyParser(p); err != nil {
		fmt.Println(err)
		return nil
	}

	*p = middleware.CheckDeviceInformation(*p)

	return p
}
