package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("César")
	expected := "Hello, César!"

	if result != expected {
		t.Errorf("The function evaluated %s, when %s was expected", result, expected)
	}

}
