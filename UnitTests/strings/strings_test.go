package strings_test

import (
	"demo/tests/strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	s := "hello"
	expected := "Hello"

	got := strings.Capitalize(s)

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
