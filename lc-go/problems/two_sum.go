package problems

// TwoSumSolution
//
// Problem link: https://leetcode.com/problems/two-sum/
func TwoSumSolution(nums []int, target int) []int {
	var solution []int

	// Compare each value with another value
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				solution = []int{i, j}
			}
		}
	}

	return solution
}

func TwoSumOptimizedSolution(nums []int, target int) []int {
	// The idea is having a reference of what next value is.
	// For instance, the nums are [2, 7, 11, 15] and the target is 9.
	// First, we create with a key of (target - element's) value:
	// {
	// 		(9-2 -> 7): 0, 0 -> index
	// 		(9-7 -> 2): 1,
	// 		(9-11 -> -2): 3,
	// 		(9-15 -> -6): 4
	// }
	//
	s := map[int]int{}

	// N
	for i := 0; i < len(nums); i++ {
		s[target-nums[i]] = i
	}

	// N
	for i := 0; i < len(nums); i++ {
		val, ok := s[nums[i]]
		if ok {
			if i == val {
				continue
			}
			if i < val {
				return []int{i, val}
			} else {
				return []int{val, i}
			}
		}

	}

	return []int{}
}
