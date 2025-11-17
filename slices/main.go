package main

import (
	"fmt"
)

// Return a new slice with the lowest score removed
// If multiple same lowest scores exist, remove only the first occurrence
func RemoveLowest(grades []int) []int {
	return nil
}

// Return the three highest scores in descending order
func GetTopThree(grades []int) []int {
	return nil
}

// Calculate average after removing the lowest score
func AverageWithoutLowest(grades []int) float64 {
	return 0.0
}

// Return the middle 60% of scores (excluding top and bottom 20%)
func GetMiddleScores(grades []int) []int {
	return nil
}

func main() {
	grades := []int{66, 92, 78, 0, 100, 88, 72, 95, 67, 83, 54}

	fmt.Println("Original grades:", grades)
	fmt.Println("After removing lowest:", RemoveLowest(grades))
	fmt.Println("Top three scores:", GetTopThree(grades))
	fmt.Printf("Average without lowest: %.2f\n", AverageWithoutLowest(grades))
	fmt.Println("Middle 60% scores:", GetMiddleScores(grades))
}
