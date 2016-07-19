package main_test

import (
	"github.com/helionce/go-env"
	"testing"
)

func TestCount(t *testing.T) {
	if main.Count() != 1 {
		t.Error("Expected 1")
	}
	if main.Count() != 2 {
		t.Error("Expected 2")
	}
	if main.Count() != 3 {
		t.Error("Expected 3")
	}
}
