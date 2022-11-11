package controllers

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/middleware"
	"hotel_cms/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PostControl(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p models.Person

		if err := c.BodyParser(&p); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(p)

		p = middleware.CheckPersonInformation(p)
		fmt.Println(p)

		q := utils.GetQueryUpdate(p)
		fmt.Println(q)
		return nil
	}
}
