package hi

import (
	"math"
	"testing"
)

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
			result := Max(tt.input)
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
	result := Max(input)
	if !result.Valid() {
		t.Errorf("Max(%v) = %t, want %t", input, result.Valid(), true)
	}
	got := result.Get()
	if !math.IsNaN(got) {
		t.Errorf("Max(%v) = %f, want NaN", input, got)
	}
}
