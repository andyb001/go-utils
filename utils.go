package main

import (
	"strconv"
	"strings"
)

// containsSubstring checks if a string contains a specified substring.
//
// Parameters:
// - mainString: the main string to search in.
// - subString: the substring to search for.
//
// Return type: bool.
func containsSubstring(mainString, subString string) bool {
	return strings.Contains(mainString, subString)
}

// stringToFloat converts a string to a float64.
//
// It takes an input string as a parameter and returns a float64 value.
// It also returns an error if the conversion fails.
func stringToFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

// stringToInt64 converts a string to an int64.
//
// It takes a string as input and returns the converted int64 value and an error if any.
func stringToInt64(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}
