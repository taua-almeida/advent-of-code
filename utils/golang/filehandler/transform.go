package filehandler

import "strings"

func TransformFileToSliceOfStrings(fileBytes []byte) []string {
	sliceOfStrings := strings.Split(string(fileBytes), "\n")
	return sliceOfStrings
}
