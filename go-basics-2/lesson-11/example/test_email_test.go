package example

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, e1 := IsEmailValid("risha@gmail.com")
	if e1 != nil {
		t.Error("Expected nil, got ", e1)
	}
	_, e2 := IsEmailValid("rishagmail.com")
	if e2 != nil {
		t.Error("Expected nil, got ", e2)
	}
	_, e3 := IsEmailValid("aryan2621@gmail.com")
	if e3 != nil {
		t.Error("Expected nil, got ", e3)
	}
}
