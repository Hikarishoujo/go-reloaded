package text_edits

import (
	"testing"
)

// TestHandleHex tests the hex number conversion
func TestHandleHex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"valid hex", "FF(hex)", "255"},
		{"invalid hex", "ZZ(hex)", "ZZ(hex)"},
		{"mixed case", "1a2B3c(hex)", "1714684"},
		{"empty", "", ""},
		{"no hex marker", "FF", "FF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleHex(tt.input); got != tt.expected {
				t.Errorf("HandleHex() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestHandleBinary tests the binary number conversion
func TestHandleBinary(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"valid binary", "1010(bin)", "10"},
		{"invalid binary", "1012(bin)", "1012(bin)"},
		{"leading zeros", "0001(bin)", "1"},
		{"empty", "", ""},
		{"no bin marker", "1010", "1010"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleBinary(tt.input); got != tt.expected {
				t.Errorf("HandleBinary() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestHandleCaseModifications tests case transformations
func TestHandleCaseModifications(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"uppercase", "Hello(up)", "HELLO"},
		{"lowercase", "Hello(low)", "hello"},
		{"capitalize", "hello(cap)", "Hello"},
		{"mixed case", "HeLLo(up)", "HELLO"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleCaseModifications(tt.input); got != tt.expected {
				t.Errorf("HandleCaseModifications() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestHandlePunctuation tests punctuation handling
func TestHandlePunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"ellipsis", "...", "..."},
		{"comma", "word, word", "word, word"},
		{"multiple", "..., word,", "..., word,"},
		{"empty", "", ""},
		{"no punctuation", "word word", "word word"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandlePunctuation(tt.input); got != tt.expected {
				t.Errorf("HandlePunctuation() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestHandleApostrophes tests apostrophe handling
func TestHandleApostrophes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single quote", "'word'", "'word'"},
		{"multiple words", "'word one'", "'word one'"},
		{"empty", "", ""},
		{"no quote", "word", "word"},
		{"multiple quotes", "'word' 'another'", "'word' 'another'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleApostrophes(tt.input); got != tt.expected {
				t.Errorf("HandleApostrophes() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestHandleArticles tests article handling
func TestHandleArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"a before vowel", "a apple", "an apple"},
		{"A before vowel", "A apple", "An apple"},
		{"a before consonant", "a cat", "a cat"},
		{"empty", "", ""},
		{"no article", "the dog", "the dog"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleArticles(tt.input); got != tt.expected {
				t.Errorf("HandleArticles() = %q, want %q", got, tt.expected)
			}
		})
	}
}
