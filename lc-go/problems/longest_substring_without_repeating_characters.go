package problems

import "fmt"

//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func LongestSubstringWithoutRepeatingChar(s string) int {
	// Sliding window
	// Temporary map track repeating character

	// only save the latest
	// Constraints:
	// 0 <= s.length <= 5 * 104
	// s consists of English letters, digits, symbols and spaces.
	count := 0
	startPointer := 0
	endPointer := 0
	// The value is the char's index.
	charOccurence := make(map[string]int)

	// While endPointer haven't reached the last element, do the loop.
	for endPointer < len(s) {
		currentChar := string([]rune(s)[endPointer])
		value, ok := charOccurence[currentChar]
		if !ok || value == -1 {
			// The current character isn't exist
			charOccurence[currentChar] = endPointer
			endPointer += 1
			if count < (endPointer - startPointer) {
				count = endPointer - startPointer
			}
		} else {
			//fmt.Println("charOccurence", charOccurence)
			//fmt.Println("currentChar", currentChar)
			startPointer = charOccurence[currentChar] + 1
			//fmt.Println("startPointer", startPointer)
			charOccurence = make(map[string]int)
			//charOccurence[currentChar] = -1
			//endPointer += 1

			//charOccurence[string([]rune(s)[startPointer])] = startPointer
		}
	}

	// start moving the endpointer
	// every movement, add the current character to the map
	return count
}

func LengthOfLongestSubstring(s string) int {
	currentLongestSubstring := ""
	longestSubstring := ""
	longestSubstringLength := 0
	lastSubstringIdx := -1
	charOccurrence := make(map[string]int)

	for idx, char := range s {
		fmt.Printf("idx: %v, char: %v, currentLongest: %v \n", idx, string(char), currentLongestSubstring)
		_, ok := charOccurrence[string(char)]
		if !ok {
			charOccurrence[string(char)] = 1
			currentLongestSubstring += string(char)
			lastSubstringIdx += 1
		} else {
			if len(currentLongestSubstring) > longestSubstringLength {
				longestSubstringLength = len(currentLongestSubstring)
				longestSubstring = currentLongestSubstring
			}
			currentLongestSubstring = string(char)
		}
	}

	//fmt.Println("current longest string:", currentLongestSubstring)
	//fmt.Println("current longest string:", longestSubstring)
	if len(longestSubstring) > len(currentLongestSubstring) {
		return len(longestSubstring)
	} else {
		return len(currentLongestSubstring)
	}
}
