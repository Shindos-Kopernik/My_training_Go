package main

import "testing"

func TestMultiple(t *testing.T) {
	var x, y, result = 2, 2, 4

	realResult := Multiple(x, y)
	if realResult != result {
		t.Errorf("%d != %d", result, realResult)
	}
}
