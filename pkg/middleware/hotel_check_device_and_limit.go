package middleware

import (
	"fmt"
	"hotel_cms/app/models"
	"math"
)

func CheckHotelInformation(p *models.HotelManage) *models.HotelManage {
	fmt.Println((*p).HotelID)
	fmt.Println(p.DevicesLimit)

	if p.DevicesLimit == nil {
		val := 0
		p.DevicesLimit = &val
	}
	if p.DevicesNumber == nil {
		val := 0
		p.DevicesNumber = &val
	}

	if *p.DevicesLimit < 0 {
		*p.DevicesLimit = 0
	}

	*p.DevicesNumber = int(math.Min(float64(*p.DevicesNumber), float64(*p.DevicesLimit)))

	return p
}
