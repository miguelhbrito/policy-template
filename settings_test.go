package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	settings := Settings{}

	if !settings.IsPalindrome("arara") {
		t.Errorf("string should be denied")
	}

	if settings.IsPalindrome("zelda") {
		t.Errorf("string should not be denied")
	}
}
