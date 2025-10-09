package twopointers

import (
	"math"
)

func MaxArea(height []int) int {
	front, back := 0, len(height)-1
	mostWater := 0
	currWater := 0

	for front < back {
		smallestHeight := int(math.Min(float64(height[front]), float64(height[back])))
		currWater = smallestHeight * (back - front)
		mostWater = int(math.Max(float64(mostWater), float64(currWater)))
		if height[front] < height[back] {
			front++
		} else {
			back--
		}
	}
	return mostWater
}
