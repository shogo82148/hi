package sets

import "testing"

func TestAll(t *testing.T) {
	testCases := []struct {
		name  string
		set   Set[int]
		value int
		want  bool
	}{
		{
			name:  "empty",
			set:   New[int](),
			value: 1,
			want:  true,
		},
		{
			name:  "one element",
			set:   New(1),
			value: 1,
			want:  true,
		},
		{
			name:  "another element",
			set:   New(2),
			value: 1,
			want:  false,
		},
		{
			name:  "multiple elements",
			set:   New(1, 2),
			value: 1,
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := All(tc.set, tc.value); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
