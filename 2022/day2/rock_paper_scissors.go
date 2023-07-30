package main

type Point uint8
type ChoiceOptions string

// Points
const (
	Loss Point = 3 * iota
	Draw
	Win
)

const (
	RockPoint Point = 1 + iota
	PaperPoint
	ScissorsPoint
)

// Game ChoiceOptions
const (
	Rock    ChoiceOptions = "Rock"
	Paper   ChoiceOptions = "Paper"
	Scissor ChoiceOptions = "Scissor"
)

// playerChoice is you and opponentChoice is the elf
func play(opponentChoice ChoiceOptions, playerChoice ChoiceOptions) Point {
	if opponentChoice == playerChoice {
		return Draw + playerChoice.choiceToPoints()
	}

	if opponentChoice == Scissor && playerChoice == Rock {
		return Win + playerChoice.choiceToPoints()
	}

	if opponentChoice == Rock && playerChoice == Paper {
		return Win + playerChoice.choiceToPoints()
	}

	if opponentChoice == Paper && playerChoice == Scissor {
		return Win + playerChoice.choiceToPoints()
	}

	return Loss + playerChoice.choiceToPoints()
}

func inputsToChoiceOptions(inputs []string) []ChoiceOptions {
	convertedChoices := make([]ChoiceOptions, 0, len(inputs))

	for _, input := range inputs {
		if input == "A" || input == "X" {
			convertedChoices = append(convertedChoices, Rock)
		}
		if input == "B" || input == "Y" {
			convertedChoices = append(convertedChoices, Paper)
		}
		if input == "C" || input == "Z" {
			convertedChoices = append(convertedChoices, Scissor)
		}
	}

	return convertedChoices
}

func (co ChoiceOptions) choiceToPoints() Point {
	switch co {
	case Rock:
		return RockPoint
	case Paper:
		return PaperPoint
	case Scissor:
		return ScissorsPoint
	}
	return 0
}

func getPlayerPoints(points []Point) int {
	sum := 0
	for _, p := range points {
		sum += int(p)
	}

	return sum
}
