package main

import "testing"

func Test_addrfunc(t *testing.T) {
	result := simpleAddr(2, 3)
	if result != 5 {
		t.Error("Incorrect result: expected 5, got", result)
	}
}
