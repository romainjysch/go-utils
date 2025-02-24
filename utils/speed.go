package utils

import (
	"fmt"
	"math"
)

func Hello() {
	fmt.Println("Hello World")
}

func PaceToSpeed(minutes int, seconds int) {
	m := float64(minutes)
	s := float64(seconds)
	speed := 3600 / (m*60 + s)
	fmt.Printf("%.2f\n", speed)
}

func SpeedToPace(speedInKmh float64) {
	t := math.Round(3600 / speedInKmh)
	minutes := int64(t) / 60
	seconds := int64(t) - 60*minutes
	fmt.Println(minutes, "minutes", seconds, "second(s) per kilometer")
}
