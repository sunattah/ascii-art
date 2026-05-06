package main

import (
	"reflect"
	"testing"
)

// TestParseInput checks if the string is split correctly by \n
func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Hello", []string{"Hello"}},
		{"Hello\\nWorld", []string{"Hello", "World"}},
		{"\\nHello", []string{"", "Hello"}},
		{"Line1\\nLine2\\nLine3", []string{"Line1", "Line2", "Line3"}},
	}

	for _, test := range tests {
		result := ParseInput(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %q, expected %v but got %v", test.input, test.expected, result)
		}
	}
}

// TestIsValid checks if the program correctly identifies allowed ASCII characters
func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ABCabc123", true},
		{"Hello World!", true},
		{"Hello\nWorld", true},
		{"Nice to meet you @", true},
		{"£50", false},      // Non-ASCII character
		{"你好", false},      // Mandarin characters
		{"é", false},         // Accented character
	}

	for _, test := range tests {
		result := IsValid(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %v but got %v", test.input, test.expected, result)
		}
	}
}