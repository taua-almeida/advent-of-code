package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestMostCalories(t *testing.T) {
	test_input := readInput("_test_elfs_food.txt")

	arrayOfCalories := makeArrayOfCalories(test_input)

	calorieCounts := generateCaloriesCount(arrayOfCalories)

	mostCalories := findMostCalories(calorieCounts)

	if mostCalories != 24000 {
		t.Errorf("Most calories is expected to be 24000, but got: %v", mostCalories)
	}
}

func TestTopThreeCalories(t *testing.T) {
	test_input := readInput("_test_elfs_food.txt")

	arrayOfCalories := makeArrayOfCalories(test_input)

	calorieCounts := generateCaloriesCount(arrayOfCalories)

	topThree := find3Largest(calorieCounts)

	if len(topThree) != 3 {
		t.Errorf("Expected to have 3 calories count, but got: %v", len(topThree))
	}

	expectedTopThree := []int{24000, 11000, 10000}

	if !slices.Equal(topThree, expectedTopThree) {
		t.Errorf("Top 3 calories expected to be: [24000, 11000, 10000], but got: %v", topThree)
	}

	totalCalories := sumSlices(topThree)

	if totalCalories != 45000 {
		t.Errorf("Sum of the top three was expected to be: 45000, but got: %v", totalCalories)
	}
}
