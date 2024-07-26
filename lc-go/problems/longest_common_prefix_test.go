package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	problem := []string{"flower", "flow", "flight"}
	answer := "fl"
	solution := LongestCommonPrefix(problem)

	if solution != answer {
		t.Errorf("invaid value. Expected %v, got %v", answer, solution)
	}

	problem = []string{"dog", "racecar", "car"}
	answer = ""
	solution = LongestCommonPrefix(problem)

	if solution != answer {
		t.Errorf("invaid value. Expected %v, got %v", answer, solution)
	}
}
