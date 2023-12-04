package main

import (
	"testing"

	"github.com/taua-almeida/advent-of-code/utils/golang/filehandler"
)

func TestAllInputValues(t *testing.T) {
	fileBytes := filehandler.ReadTxtFile("../input.txt")
	sliceOfStrings := filehandler.TransformFileToSliceOfStrings(fileBytes)
	if len(sliceOfStrings) != 1000 {
		t.Errorf("Expected 1000, got %v", len(sliceOfStrings))
	}
}

func TestInputValuesTest(t *testing.T) {
	fileBytes := filehandler.ReadTxtFile("../input_test_pt1.txt")
	sliceOfStrings := filehandler.TransformFileToSliceOfStrings(fileBytes)
	if len(sliceOfStrings) != 4 {
		t.Errorf("Expected 4, got %v", len(sliceOfStrings))
	}
	if sliceOfStrings[0] != "1abc2" {
		t.Errorf("First string expected 1abc2, got %v", sliceOfStrings[0])
	}
}

func TestCalibrationCalculationPt1(t *testing.T) {
	result := SolvePart1()
	if result != 142 {
		t.Errorf("Expected 142, got %v", result)
	}
}

func TestCalibrationCalculationPt2(t *testing.T) {
	result := SolvePart2()
	if result != 281 {
		t.Errorf("Expected 281, got %v", result)
	}
}
