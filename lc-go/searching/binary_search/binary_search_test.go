package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Print(5 / 2)

	arr := []int{1, 2, 3, 5, 10}
	//arrP := &arr

	//fmt.Println((*arrP)[50])
	result := binarySearch(arr, 10)
	if result != 2 {
		t.Errorf("invalid result")
	}
}
