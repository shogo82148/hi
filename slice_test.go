package hi

import (
	"math"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		target   float64
		expected int
	}{
		{
			name:     "empty slice",
			input:    []float64{},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element slice",
			input:    []float64{1},
			target:   1,
			expected: 1,
		},
		{
			name:     "multiple element slice",
			input:    []float64{1, 2, 3, 4, 3, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "NaN",
			input:    []float64{0, math.NaN(), 1, math.NaN(), 2},
			target:   math.NaN(),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Count(tt.input, tt.target)
			if result != tt.expected {
				t.Errorf("Count(%v, %f) = %d, want %d", tt.input, tt.target, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		valid    bool
		expected float64
	}{
		{
			name:  "empty slice",
			input: []float64{},
			valid: false,
		},
		{
			name:     "single element slice",
			input:    []float64{1},
			valid:    true,
			expected: 1,
		},
		{
			name:     "multiple element slice",
			input:    []float64{1, 2, 3, 4, 5},
			valid:    true,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.input...)
			if result.Valid() != tt.valid {
				t.Errorf("Max(%v) = %t, want %t", tt.input, result.Valid(), tt.valid)
			}
			if result.Valid() {
				v := result.Get()
				if v != tt.expected {
					t.Errorf("Max(%v) = %f, want %f", tt.input, v, tt.expected)
				}
			}
		})
	}
}

// Test for slices includes NaN.
func TestMax_Nan(t *testing.T) {
	input := []float64{
		0, 1, 2, math.NaN(), 4, 5,
	}
	result := Max(input...)
	if !result.Valid() {
		t.Errorf("Max(%v) = %t, want %t", input, result.Valid(), true)
	}
	got := result.Get()
	if !math.IsNaN(got) {
		t.Errorf("Max(%v) = %f, want NaN", input, got)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		valid    bool
		expected float64
	}{
		{
			name:  "empty slice",
			input: []float64{},
			valid: false,
		},
		{
			name:     "single element slice",
			input:    []float64{1},
			valid:    true,
			expected: 1,
		},
		{
			name:     "multiple element slice",
			input:    []float64{1, 2, 3, 4, 5},
			valid:    true,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.input...)
			if result.Valid() != tt.valid {
				t.Errorf("Min(%v) = %t, want %t", tt.input, result.Valid(), tt.valid)
			}
			if result.Valid() {
				v := result.Get()
				if v != tt.expected {
					t.Errorf("Min(%v) = %f, want %f", tt.input, v, tt.expected)
				}
			}
		})
	}
}

// Test for slices includes NaN.
func TestMin_Nan(t *testing.T) {
	input := []float64{
		0, 1, 2, math.NaN(), 4, 5,
	}
	result := Min(input...)
	if !result.Valid() {
		t.Errorf("Max(%v) = %t, want %t", input, result.Valid(), true)
	}
	got := result.Get()
	if !math.IsNaN(got) {
		t.Errorf("Max(%v) = %f, want NaN", input, got)
	}
}
