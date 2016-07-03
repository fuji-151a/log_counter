package main

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	path := "../../data/"
	actual := listFiles(path, path)
	expected := 15
	if actual != expected {
		t.Errorf("expected is %d. but found actual is %d", expected, actual)
	}
}
