package problems

import (
	"reflect"
	"testing"
)

func TestThreeSumSolution(t *testing.T) {
	// ====================
	// Case 1
	// ====================
	nums := []int{-1, 0, 1, 2, -1, -4}
	//target := 0
	answer := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	solution := ThreeSumSolution(nums)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 2
	// ====================
	nums = []int{0, 1, 1}
	//target = 0
	answer = [][]int{}
	solution = ThreeSumSolution(nums)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 3
	// ====================
	nums = []int{0, 0, 0}
	//target = 0
	answer = [][]int{{0, 0, 0}}
	solution = ThreeSumSolution(nums)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}
}
