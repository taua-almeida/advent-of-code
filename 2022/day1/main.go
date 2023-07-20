package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	foodInput, err := ioutil.ReadFile("elfs_food.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	arrayOfCalories := strings.Split(string(foodInput), "\n")
	calorieCounts := make([]int, 0)

	count := 0
	for _, cal := range arrayOfCalories {
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

	mostCalories := 0
	for _, cal := range calorieCounts {
		if cal > mostCalories {
			mostCalories = cal
		}
	}

	fmt.Println("Most calories is:", mostCalories)
}
