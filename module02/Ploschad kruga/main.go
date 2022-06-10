package main

import (
	"fmt"
	"math"
)

func main() {
	const Pi = 3.14
	var radius float64
	var long float64 = 35
	var r *float64 // r обявляем как указатель на float64
	r = &radius

	radius = long / (2 * Pi)                  //находим радиус круга
	fmt.Println(math.Round(radius*100) / 100) // печатаем радиус круга (округленную)

	s := Pi * (*r * *r)                  //находим площу круга
	fmt.Println(math.Round(s*100) / 100) // печатаем площу круга (округленную)
}
