package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	q := 121
	a := true
	s := IsPalindrome(q)

	if a != s {
		t.Errorf("invalid value. Expected %v, got %v", a, s)
	}

	q = -121
	a = false
	s = IsPalindrome(q)

	if a != s {
		t.Errorf("invalid value. Expected %v, got %v", a, s)
	}

	q = 10
	a = false
	s = IsPalindrome(q)

	if a != s {
		t.Errorf("invalid value. Expected %v, got %v", a, s)
	}

	q = 3
	a = true
	s = IsPalindrome(q)

	if a != s {
		t.Errorf("invalid value. Expected %v, got %v", a, s)
	}

	q = 443
	a = false
	s = IsPalindrome(q)

	if a != s {
		t.Errorf("invalid value. Expected %v, got %v", a, s)
	}
}
