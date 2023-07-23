package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func sumSlices(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func makeArrayOfCalories(caloriesString []byte) []string {
	arrayOfCalories := strings.Split(string(caloriesString), "\n")
	return arrayOfCalories
}

func readInput(name string) []byte {
	foodInput, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return foodInput
}

func findMostCalories(calories []int) int {
	mostCalories := 0
	for _, cal := range calories {
		if cal > mostCalories {
			mostCalories = cal
		}
	}
	return mostCalories
}

func find3Largest(calories []int) []int {
	threeLargest := []int{0, 0, 0}

	first := math.MinInt
	second := math.MinInt
	third := math.MinInt

	for _, cal := range calories {
		if cal > first {
			third = second
			second = first
			first = cal
		} else if cal > second {
			third = second
			second = cal
		} else if cal > third {
			third = cal
		}
	}

	threeLargest[0] = first
	threeLargest[1] = second
	threeLargest[2] = third
	return threeLargest
}

func generateCaloriesCount(calories []string) []int {
	calorieCounts := make([]int, 0)

	count := 0
	for _, cal := range calories {
		if cal == "" {
			calorieCounts = append(calorieCounts, count)
			count = 0
			continue
		}

		calInt, err := strconv.Atoi(cal)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}
		count += calInt
	}
	calorieCounts = append(calorieCounts, count)

	return calorieCounts
}

func main() {

	input := readInput("elfs_food.txt")
	arrayOfCalories := makeArrayOfCalories(input)

	calorieCounts := generateCaloriesCount(arrayOfCalories)

	fmt.Println("Most calories: ", findMostCalories(calorieCounts))

	topThree := find3Largest(calorieCounts)
	fmt.Println("Top three calories: ", topThree)

	fmt.Println("Sum of top three calories: ", sumSlices(topThree))

}
