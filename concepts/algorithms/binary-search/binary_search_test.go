package binarysearch

import "testing"

func TestBounds(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		target    int
		wantLower int
		wantUpper int
	}{
		{name: "empty", values: nil, target: 3, wantLower: 0, wantUpper: 0},
		{name: "before all", values: []int{2, 4, 6}, target: 1, wantLower: 0, wantUpper: 0},
		{name: "duplicate range", values: []int{1, 3, 3, 3, 5}, target: 3, wantLower: 1, wantUpper: 4},
		{name: "between values", values: []int{1, 3, 5}, target: 4, wantLower: 2, wantUpper: 2},
		{name: "after all", values: []int{1, 3, 5}, target: 8, wantLower: 3, wantUpper: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LowerBound(test.values, test.target); got != test.wantLower {
				t.Errorf("LowerBound() = %d, want %d", got, test.wantLower)
			}
			if got := UpperBound(test.values, test.target); got != test.wantUpper {
				t.Errorf("UpperBound() = %d, want %d", got, test.wantUpper)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		target int
		want   int
	}{
		{name: "finds first duplicate", values: []int{1, 2, 2, 4}, target: 2, want: 1},
		{name: "missing", values: []int{1, 2, 4}, target: 3, want: -1},
		{name: "empty", values: nil, target: 3, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Search(test.values, test.target); got != test.want {
				t.Fatalf("Search() = %d, want %d", got, test.want)
			}
		})
	}
}
