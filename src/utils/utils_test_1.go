package utils

import "testing"

func TestSuccessStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b", "c"}); ok {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}

func TestFailedStringInSlice(t *testing.T) {
	if ok := StringInSlice("d", []string{"a", "b", "c"}); ok {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}
