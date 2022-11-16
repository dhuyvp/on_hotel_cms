package middleware

import (
	"hotel_cms/app/models"
	"math"
)

func CheckPersonInformation(p models.Person) models.Person {
	if *p.Age < 0 {
		*p.Age = 0
	}

	*p.Age = int(math.Min(float64(*p.Age), float64(10)))

	return p
}
