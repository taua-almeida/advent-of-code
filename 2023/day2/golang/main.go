package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/taua-almeida/advent-of-code/utils/golang/filehandler"
)

type CubeCounts struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID     int
	Counts []CubeCounts
}

// Used for part1
const (
	MaxRedCubes   = 12
	MaxGreenCubes = 13
	MaxBlueCubes  = 14
)

func main() {
	fileBytes := filehandler.ReadTxtFile("../input.txt")

	sliceOfStrings := filehandler.TransformFileToSliceOfStrings(fileBytes)

	games, err := parseGames(sliceOfStrings)
	if err != nil {
		fmt.Println("Error parsing games:", err)
		return
	}

	gameIdSum := 0
	for _, game := range games {
		if isPlayable(game) {
			gameIdSum += game.ID
		}
	}
	fmt.Println("Sum of playable games IDs:", gameIdSum)

	powerSetSum := 0
	for _, game := range games {
		powerSetSum += getHighestCubeCount(game)
	}
	fmt.Println("Sum of power set:", powerSetSum)
}

func parseGames(gameStrings []string) ([]*Game, error) {
	games := make([]*Game, 0, len(gameStrings))

	for _, line := range gameStrings {
		parts := strings.Split(line, ":")
		gameID, err := strconv.Atoi(strings.TrimSpace(parts[0][4:]))
		if err != nil {
			return nil, fmt.Errorf("error converting game ID to int: %w", err)
		}

		game := &Game{ID: gameID}

		subsets := strings.Split(parts[1], ";")
		for _, subset := range subsets {
			counts, err := parseCubeCounts(subset)
			if err != nil {
				return nil, err
			}
			game.Counts = append(game.Counts, *counts)
		}
		games = append(games, game)

	}

	return games, nil
}

func parseCubeCounts(subset string) (*CubeCounts, error) {
	cubes := strings.Split(subset, ",")
	counts := CubeCounts{}
	for _, cube := range cubes {
		parts := strings.Split(strings.TrimSpace(cube), " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid cube data: %s", cube)
		}
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error converting cube count to int: %w", err)
		}
		switch parts[1] {
		case "red":
			counts.Red += count
		case "green":
			counts.Green += count
		case "blue":
			counts.Blue += count
		default:
			return nil, fmt.Errorf("unknown cube color: %s", parts[1])
		}
	}
	return &counts, nil
}

// Used only for pt1
func isPlayable(game *Game) bool {
	for _, counts := range game.Counts {
		if counts.Red > MaxRedCubes || counts.Green > MaxGreenCubes || counts.Blue > MaxBlueCubes {
			return false
		}
	}
	return true
}

// Used only for pt2
func getHighestCubeCount(game *Game) int {
	highestHighestRed := 0
	highestHighestGreen := 0
	highestHighestBlue := 0
	for _, counts := range game.Counts {
		if counts.Red > highestHighestRed {
			highestHighestRed = counts.Red
		}
		if counts.Green > highestHighestGreen {
			highestHighestGreen = counts.Green
		}
		if counts.Blue > highestHighestBlue {
			highestHighestBlue = counts.Blue
		}
	}
	return highestHighestRed * highestHighestGreen * highestHighestBlue
}
