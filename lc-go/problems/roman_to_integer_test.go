package problems

import (
	"fmt"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	s := "III"
	answer := 3
	solution := RomanToInt(s)

	fmt.Println("solution ", solution)
	fmt.Println("answer", answer)
	if solution != answer {
		t.Errorf("invalid value for %v! Expected %v, got %v", s, answer, solution)
	}

	s = "IV"
	answer = 4
	solution = RomanToInt(s)
	fmt.Println("solution ", solution)
	fmt.Println("answer", answer)
	if solution != answer {
		t.Errorf("invalid value for %v! Expected %v, got %v", s, answer, solution)
	}

	s = "MCMXCIV"
	answer = 1994
	solution = RomanToInt(s)
	fmt.Println("solution ", solution)
	fmt.Println("answer", answer)
	if solution != answer {
		t.Errorf("invalid value for %v! Expected %v, got %v", s, answer, solution)
	}

	s = "LVIII"
	answer = 58
	solution = RomanToInt(s)
	fmt.Println("solution ", solution)
	fmt.Println("answer", answer)
	if solution != answer {
		t.Errorf("invalid value for %v! Expected %v, got %v", s, answer, solution)
	}
}
