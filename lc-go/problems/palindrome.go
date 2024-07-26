package problems

import (
	"strconv"
)

func IsPalindrome(x int) bool {

	isPalindrome := true
	stringOfX := strconv.Itoa(x)

	runeOfX := []rune(stringOfX)
	lengthOfX := len(stringOfX)

	for i := 0; i < (lengthOfX / 2); i++ {
		if runeOfX[(lengthOfX-1)-i] != runeOfX[i] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
