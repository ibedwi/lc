package problems

import (
	"fmt"
	"math"
)

// Problem link: https://leetcode.com/problems/longest-common-prefix/submissions/1023342065/
func LongestCommonPrefix(strs []string) string {
	// Compared Value.
	// - If the value is "", then stop the iteration early and return empty string
	// - Else, use as compared value. Substract until we have a common value
	// Start from first element
	// Compare a and b, where "a" is the prev element and b is the current element

	if len(strs) < 2 {
		return strs[0]
	}

	// First comparation
	commonString := getCommonString(strs[0], strs[1])

	for i := 1; i < len(strs); i++ {
		commonString = getCommonString(commonString, strs[i])
		fmt.Println("iteration: ", i, " | common string", commonString)
		if commonString == "" {
			break
		}
	}

	return commonString
}

func getCommonString(str1 string, str2 string) string {
	commonString := ""
	str1Rune := []rune(str1)
	str2Rune := []rune(str2)

	maxIteration := math.Min(float64(len(str1)), float64(len(str2)))

	for i := 0; i < int(maxIteration); i++ {
		if string(str1Rune[i]) != string(str2Rune[i]) {
			break
		}

		commonString += string(str1Rune[i])
	}

	return commonString
}
