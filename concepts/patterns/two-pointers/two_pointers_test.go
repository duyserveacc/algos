package twopointers

import "testing"

func TestPairSumSorted(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		target    int
		wantLeft  int
		wantRight int
		wantOK    bool
	}{
		{name: "finds pair", values: []int{1, 2, 4, 7, 11}, target: 9, wantLeft: 1, wantRight: 3, wantOK: true},
		{name: "supports negatives", values: []int{-8, -2, 3, 9}, target: 7, wantLeft: 1, wantRight: 3, wantOK: true},
		{name: "requires distinct positions", values: []int{4}, target: 8, wantOK: false},
		{name: "missing pair", values: []int{1, 2, 3}, target: 20, wantOK: false},
		{name: "empty", values: nil, target: 0, wantOK: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			left, right, ok := PairSumSorted(test.values, test.target)
			if ok != test.wantOK {
				t.Fatalf("PairSumSorted() ok = %t, want %t", ok, test.wantOK)
			}
			if ok && (left != test.wantLeft || right != test.wantRight) {
				t.Fatalf("PairSumSorted() = (%d, %d), want (%d, %d)", left, right, test.wantLeft, test.wantRight)
			}
		})
	}
}
