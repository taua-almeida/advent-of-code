package main

import (
	"log"
	"os"
	"strings"
)

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	return file
}

func splitInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func playerChoicesAndOpponentChoices(inputs []string) ([]ChoiceOptions, []ChoiceOptions) {
	opponentChoices := make([]string, 0)
	playerChoices := make([]string, 0)

	for _, choice := range inputs {
		abChoice := strings.Split(choice, " ")
		opponentChoices = append(opponentChoices, abChoice[0])
		playerChoices = append(playerChoices, abChoice[1])
	}

	return inputsToChoiceOptions(opponentChoices), inputsToChoiceOptions(playerChoices)
}
