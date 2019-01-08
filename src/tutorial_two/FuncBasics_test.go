package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Add was incorrect, got :%d, want %d.", total, 10)
	}
}
