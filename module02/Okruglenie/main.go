package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	bike1 := 120.4 // скорость в м/с
	bike2 := 130.0 // скорость в м/с

	const km = 1000.0
	const mile = 1609.0

	speed_bike1 := AmericanVelocity(convert(bike1, km))
	speed_bike2 := EuropeanVelocity(convert(bike2, mile))

	fmt.Printf("Байк номер 1 едет %v км/час \n", speed_bike1)
	fmt.Printf("Байк номер 2 едет %v миль/час \n", speed_bike2)
}

// функция конвертации м/с в км/час   или   м/с в мили/час (результат округляется к двум знакам после запятой)
func convert(s float64, v float64) float64 {
	speed := (s * 3600) / v
	result_speed := math.Round(speed*100) / 100
	return result_speed
}
