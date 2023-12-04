package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/taua-almeida/advent-of-code/utils/golang/filehandler"
)

func calculateCalibrationPt1(texts []string) int {
	var result int

	for _, value := range texts {
		var firstDigit rune
		var lastDigit rune

		// Find the first digit
		for _, char := range value {
			if char >= '0' && char <= '9' {
				firstDigit = char
				break
			}
		}

		// Find the last digit
		for i := len(value) - 1; i >= 0; i-- {
			if value[i] >= '0' && value[i] <= '9' {
				lastDigit = rune(value[i])
				break
			}
		}

		if firstDigit != 0 && lastDigit != 0 {
			digitString := string(firstDigit) + string(lastDigit)
			intDigit, err := strconv.Atoi(digitString)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			result += intDigit
		}
	}

	return result
}

func calculateCalibrationPt2(texts []string) int {
	var result int = 0
	digitMap := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	first := regexp.MustCompile("(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)")
	last := regexp.MustCompile(".*(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)")

	for _, value := range texts {
		firstDigit := first.FindStringSubmatch(value)
		lastDigit := last.FindStringSubmatch(value)

		firstValue := 10 * (slices.Index(digitMap, firstDigit[1])%9 + 1)
		secondValue := slices.Index(digitMap, lastDigit[1])%9 + 1

		result += firstValue + secondValue
	}

	return result
}

func SolvePart1() int {
	fileBytesPt1 := filehandler.ReadTxtFile("../input_test_pt1.txt")
	textsPt1 := filehandler.TransformFileToSliceOfStrings(fileBytesPt1)
	result_pt1 := calculateCalibrationPt1(textsPt1)
	fmt.Println("Result Part1: ", result_pt1)
	return result_pt1
}

func SolvePart2() int {
	fileBytesPt2 := filehandler.ReadTxtFile("../input_test_pt2.txt")
	textsPt2 := filehandler.TransformFileToSliceOfStrings(fileBytesPt2)
	result_pt2 := calculateCalibrationPt2(textsPt2)
	fmt.Println("Result Part2: ", result_pt2)
	return result_pt2
}

func main() {

	//SolvePart1()

	SolvePart2()
}
