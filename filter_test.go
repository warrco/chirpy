package main

import "testing"

func TestWords(t *testing.T) {
	result := profanityFilter("I am a fornax")
	if result != "I am a ****" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result, "I am a ****")
	}
}

func TestWords2(t *testing.T) {
	result := profanityFilter("I am a FORnax?")
	if result != "I am a FORnax?" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result, "I am a FORnax?")
	}
}

func TestWords3(t *testing.T) {
	result := profanityFilter("You're quite the KerFufflE")
	if result != "You're quite the ****" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result, "You're quite the ****")
	}
}
