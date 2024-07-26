package problems

// RomanToInt
//
// Problem link: https://leetcode.com/problems/roman-to-integer/
// ====================
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// ====================
func RomanToInt(s string) int {
	// Reference table
	// Always check next character

	solution := 0
	sRune := []rune(s)

	for i := 0; i < len(sRune); i++ {

		// If it's the last character, just add to the solution
		if i == len(s)-1 {
			value := mapCharacter(sRune[i])
			solution += value
		} else {
			if isSubstract(string(sRune[i]), string(sRune[i+1])) {
				solution += mapCharacter(sRune[i+1]) - mapCharacter(sRune[i])
				i += 1
			} else {
				solution += mapCharacter(sRune[i])
			}
		}

	}

	return solution
}

func isSubstract(current string, next string) bool {
	switch current {
	case "I":
		if next == "X" || next == "V" {
			return true
		}
	case "X":
		if next == "L" || next == "C" {
			return true
		}

	case "C":
		if next == "D" || next == "M" {
			return true
		}
	default:
		return false
	}
	return false
}

func mapCharacter(s rune) int {
	// ====================
	// Symbol       Value
	// I             1
	// V             5
	// X             10
	// L             50
	// C             100
	// D             500
	// M             1000
	// ====================
	switch string(s) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
