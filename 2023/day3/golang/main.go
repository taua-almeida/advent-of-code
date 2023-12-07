package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/taua-almeida/advent-of-code/utils/golang/filehandler"
)

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func sumPartNumbers(schematic []string) int {
	height := len(schematic)
	width := len(schematic[0])
	var sum int
	processed := make([][]bool, height)
	for i := range processed {
		processed[i] = make([]bool, width)
	}

	for y, row := range schematic {
		for x, char := range row {
			if isDigit(char) && !processed[y][x] {
				// Identify the whole number
				numStr := string(char)
				processed[y][x] = true
				j := x + 1
				for j < width && isDigit(rune(schematic[y][j])) {
					numStr += string(schematic[y][j])
					processed[y][j] = true
					j++
				}

				num, _ := strconv.Atoi(numStr)
				// Check if this number or any part of it is adjacent to a symbol
				isAdjacent := false
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						for k := 0; k < len(numStr); k++ {
							ny, nx := y+dy, x+k+dx
							if ny >= 0 && ny < height && nx >= 0 && nx < width && isSymbol(rune(schematic[ny][nx])) {
								isAdjacent = true
								break
							}
						}
						if isAdjacent {
							break
						}
					}
					if isAdjacent {
						break
					}
				}

				if isAdjacent {
					sum += num
				}
			}
		}
	}

	return sum
}

func main() {
	fileBytes := filehandler.ReadTxtFile("../input.txt")
	schematic := filehandler.TransformFileToSliceOfStrings(fileBytes)

	fmt.Println("Sum of part numbers:", sumPartNumbers(schematic))
}
