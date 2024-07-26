package problems

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	// ====================
	// Case 1
	// ====================
	nums := []int{2, 7, 11, 15}
	target := 9
	answer := []int{0, 1}
	solution := TwoSumSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 2
	// ====================
	nums = []int{3, 2, 4}
	target = 6
	answer = []int{1, 2}
	solution = TwoSumSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 3
	// ====================
	nums = []int{3, 3}
	target = 6
	answer = []int{0, 1}
	solution = TwoSumSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}
}

func TestTwoSumOptimizedSolution(t *testing.T) {
	// ====================
	// Case 1
	// ====================
	nums := []int{2, 7, 11, 15}
	target := 9
	answer := []int{0, 1}
	solution := TwoSumOptimizedSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 2
	// ====================
	nums = []int{3, 2, 4}
	target = 6
	answer = []int{1, 2}
	solution = TwoSumOptimizedSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 3
	// ====================
	nums = []int{3, 3}
	target = 6
	answer = []int{0, 1}
	solution = TwoSumOptimizedSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}

	// ====================
	// Case 4
	// ====================
	nums = []int{-1, -2, -3, -4, -5}
	target = -8
	answer = []int{2, 4}
	solution = TwoSumOptimizedSolution(nums, target)

	if !reflect.DeepEqual(solution, answer) {
		t.Errorf("Error: expected %v, got %v", answer, solution)
	}
}
