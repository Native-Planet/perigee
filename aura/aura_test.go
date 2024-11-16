package aura

import (
	"math/big"
	"testing"

	"github.com/stevelacy/go-urbit/noun"
)

func TestCord2Uw(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", "0w0"},
		{"a", "a", "0w1x"},
		{"b", "b", "0w1y"},
		{"reid", "reid", "0w1.AqmlO"},
		{"reid1", "reid1", "0w35.AqmlO"},
		{"reid12", "reid12", "0wcz5.AqmlO"},
		{"reid123", "reid123", "0wPcz5.AqmlO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cord2Uw(tt.input)
			if got != tt.want {
				t.Errorf("Cord2Uw(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestUw2Cord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "0w0", ""},
		{"a", "0w1x", "a"},
		{"b", "0w1y", "b"},
		{"reid", "0w1.AqmlO", "reid"},
		{"reid1", "0w35.AqmlO", "reid1"},
		{"reid12", "0wcz5.AqmlO", "reid12"},
		{"reid123", "0wPcz5.AqmlO", "reid123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uw2Cord(tt.input)
			if err != nil {
				t.Errorf("Uw2Cord(%q) error = %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Uw2Cord(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestUint642Uw(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  string
	}{
		{"zero", 0, "0w0"},
		{"one", 1, "0w1"},
		{"small", 42, "0wG"},
		{"big", 1234567890, "0w1.9Bwbi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint642Uw(tt.input)
			if got != tt.want {
				t.Errorf("Uint642Uw(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestUw2Uint64(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  uint64
	}{
		{"zero", "0w0", 0},
		{"one", "0w1", 1},
		{"small", "0wG", 42},
		{"big", "0w1.9Bwbi", 1234567890},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uw2Uint64(tt.input)
			if err != nil {
				t.Errorf("Uw2Uint64(%q) error = %v", tt.input, err)
				return
			}
			if got != tt.want {
				t.Errorf("Uw2Uint64(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	inputs := []string{
		"",
		"a",
		"b",
		"reid",
		"reid1",
		"reid12",
		"reid123",
		"reid1234",
		"reid12345",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			encoded := Cord2Uw(input)
			decoded, err := Uw2Cord(encoded)
			if err != nil {
				t.Errorf("Round trip error for %q: %v", input, err)
				return
			}
			if decoded != input {
				t.Errorf("Round trip failed for %q: got %q", input, decoded)
			}
		})
	}
}

func TestAtom2Uw(t *testing.T) {
	tests := []struct {
		name  string
		input *big.Int
		want  string
	}{
		{"zero", big.NewInt(0), "0w0"},
		{"one", big.NewInt(1), "0w1"},
		{"atom42", big.NewInt(42), "0wG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atom := noun.Atom{Value: tt.input}
			got := Atom2Uw(atom)
			if got != tt.want {
				t.Errorf("Atom2Uw(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestInvalidUw(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"no prefix", "123"},
		{"wrong prefix", "1w123"},
		{"invalid char", "0w123$"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Uw2Cord(tt.input)
			if err == nil {
				t.Errorf("Uw2Cord(%q) should return error", tt.input)
			}
		})
	}
}
