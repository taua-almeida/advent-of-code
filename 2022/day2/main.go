package main

import "fmt"

func main() {
	input := readFile("puzzle_input.txt")

	opponentChoices, playerChoices := playerChoicesAndOpponentChoices(splitInput(input))
	playerPoints := make([]Point, 0)

	fmt.Println(len(opponentChoices), len(playerChoices))

	for i := 0; i < len(playerChoices); i++ {
		playerPoints = append(playerPoints, play(opponentChoices[i], playerChoices[i]))
	}
	fmt.Println(len(playerPoints))

	fmt.Println(getPlayerPoints(playerPoints))
}
