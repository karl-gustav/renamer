package main

import "testing"

func TestGenerateNewName(t *testing.T) {
	newName := generateNewName("file@@@.txt", 1)
	if newName != "file001.txt" {
		t.Errorf("generateNewName() failed: %s", newName)
	}
}
