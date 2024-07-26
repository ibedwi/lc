package binary_search

import "fmt"

// binarySearch
//
// `arr` is sorted array
func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	start := 0
	end := len(arr)
	var current int
	//s := a
	for start != end {
		current = (end - start) / 2
		fmt.Printf("Start %v, Current: %v, End: %v\n", start, end, current)
		if arr[current] == target {
			return current
		} else {
			if target > arr[current] {
				start = current
			} else {
				end = current
			}

			fmt.Printf("NEW:: Start %v, Current: %v, End: %v\n", start, current, end)
		}
	}

	// Not found
	return -1
}

// 0 -- 5,
