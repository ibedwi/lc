package valid_parentheses

import (
	"fmt"
	"reflect"
	"testing"
)

type testTuple struct {
	problem string
	answer  bool
}

func TestParentheses(t *testing.T) {
	// Case 1
	// s := "()"
	// answer := true

	arr := []testTuple{
		{"(", false},
		{"()", true},
		{"(){}[]", true},
		{"){", false},
		{"()}{}", false},
		{"(){}}{", false},
		{"({[()[]{}]})", true},
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println("String: ", arr[i].problem)
		solution := IsValid(arr[i].problem)
		if !reflect.DeepEqual(arr[i].answer, solution) {
			t.Errorf("Error: expected %v, got %v", arr[i].answer, solution)
		}
	}

}
