package middleware

import (
	"hotel_cms/app/models"
	"math"
)

func CheckHotelInformation(p models.HotelManage) models.HotelManage {
	if *p.DevicesLimit < 0 {
		*p.DevicesLimit = 0
	}

	*p.DevicesNumber = int(math.Min(float64(*p.DevicesNumber), float64(*p.DevicesLimit)))

	return p
}

func CheckPersonInformation(p models.Person) models.Person {
	if *p.Age < 0 {
		*p.Age = 0
	}

	*p.Age = int(math.Min(float64(*p.Age), float64(10)))

	return p
}
