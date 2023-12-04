package main

import "testing"

func TestParseCubeCounts(t *testing.T) {
	// Test case with valid input
	input := "3 red, 2 green, 1 blue"
	expected := CubeCounts{Red: 3, Green: 2, Blue: 1}
	result, err := parseCubeCounts(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *result != expected {
		t.Errorf("Expected %v, got %v", expected, *result)
	}

	// Test case with invalid input
	invalidInput := "3 red, xyz"
	_, err = parseCubeCounts(invalidInput)
	if err == nil {
		t.Error("Expected an error for invalid input, got nil")
	}
}

func TestParseGames(t *testing.T) {
	// Prepare input data
	gameStrings := []string{"Game 1: 3 red, 2 green; 4 blue, 1 green", "Game 2: 5 red, 3 blue"}

	// Call the function
	games, err := parseGames(gameStrings)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Validate the results
	if len(games) != 2 {
		t.Errorf("Expected 2 games, got %d", len(games))
	}
}

func TestIsPlayable(t *testing.T) {
	// Test case with a playable game
	playableGame := &Game{
		ID: 1,
		Counts: []CubeCounts{
			{Red: 5, Green: 5, Blue: 5},
		},
	}
	if !isPlayable(playableGame) {
		t.Error("Expected the game to be playable")
	}

	// Test case with a non-playable game
	nonPlayableGame := &Game{
		ID: 2,
		Counts: []CubeCounts{
			{Red: 13, Green: 5, Blue: 5}, // Exceeds MaxRedCubes
		},
	}
	if isPlayable(nonPlayableGame) {
		t.Error("Expected the game to be non-playable")
	}
}
