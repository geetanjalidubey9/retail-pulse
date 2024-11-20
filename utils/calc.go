package utils

import (
	"math/rand"
	"time"
)

func CalculatePerimeter(width, height int) int {
	return 2 * (width + height)
}

func SimulateProcessingTime() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
}
