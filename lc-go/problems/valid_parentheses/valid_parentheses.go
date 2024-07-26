package valid_parentheses

// https://leetcode.com/problems/valid-parentheses/description/
func IsValid(s string) bool {
	// Check the length
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	// create stack
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// If the element is opening parentheses,
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		// The current element must be closing parentheses.
		if i == 0 || len(stack) == 0 {
			return false
		}

		lastElement := stack[len(stack)-1]
		if (lastElement == '(' && s[i] == ')') ||
			(lastElement == '{' && s[i] == '}') ||
			(lastElement == '[' && s[i] == ']') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}
