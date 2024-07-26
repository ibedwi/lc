package problems

import "fmt"

func twoSum(nums []int, target int, skipIndex int) []int {
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
		if i == skipIndex {
			continue
		}
		s[target-nums[i]] = i
	}

	// N
	for i := 0; i < len(nums); i++ {
		if i == skipIndex {
			continue
		}
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

func ThreeSumSolution(nums []int) [][]int {
	//s := make(map[int]int)

	// Extended solution from two sum problem
	// []int{-1, 0, 1, 2, -1, -4}
	// {
	// 		(t - s[i]): i
	// }
	// Then, do the same thing
	// 0 = a + b + c
	// -c = (a + b)

	//{
	//	(b): i_a
	//}

	s1 := map[int][]int{}
	var solution [][]int

	elementIsUsed := map[int]bool{}

	// N
	for i := 0; i < len(nums); i++ {
		//fmt.Printf("target: %v, i: %v, nums[i]: %v, (target - nums[i]): %v\n", target, i, nums[i], target-nums[i])
		s1[i] = twoSum(nums, 0-nums[i], i)
	}

	// pop out each element
	//fmt.Printf("================ \n")
	for i := 0; i < len(nums); i++ {
		val, ok := s1[i]
		if !ok {
			continue
		}

		if len(val) == 0 {
			continue
		}

		if elementIsUsed[i] ||
			elementIsUsed[val[0]] ||
			elementIsUsed[val[1]] {
			continue
		}

		fmt.Printf("i: %v, val: %v | nums[i]: %v, nums[val[0]]: %v, nums[val[1]]: %v \n", i, val, nums[i], nums[val[0]], nums[val[1]])
		fmt.Printf("elementIsUsed: %v", elementIsUsed)
		solution = append(solution, []int{nums[i], nums[val[0]], nums[val[1]]})
		elementIsUsed[i] = true
		elementIsUsed[val[0]] = true
		elementIsUsed[val[1]] = true
		// if the current s1 is not null
		// and the current index is not used

	}

	fmt.Printf("nums %v\n", nums)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("solution: %v\n", solution)

	return solution
}
