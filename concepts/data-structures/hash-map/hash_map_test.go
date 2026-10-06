package hashmap

import "testing"

func TestPairSumIndices(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		target     int
		wantFirst  int
		wantSecond int
		wantOK     bool
	}{
		{
			name:       "complement seen earlier",
			values:     []int{2, 7, 11, 15},
			target:     9,
			wantFirst:  0,
			wantSecond: 1,
			wantOK:     true,
		},
		{
			name:       "equal values use different indices",
			values:     []int{3, 3},
			target:     6,
			wantFirst:  0,
			wantSecond: 1,
			wantOK:     true,
		},
		{
			name:   "no pair",
			values: []int{1, 2, 4},
			target: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			first, second, ok := PairSumIndices(test.values, test.target)
			if first != test.wantFirst || second != test.wantSecond || ok != test.wantOK {
				t.Fatalf(
					"PairSumIndices() = (%d, %d, %t), want (%d, %d, %t)",
					first,
					second,
					ok,
					test.wantFirst,
					test.wantSecond,
					test.wantOK,
				)
			}
		})
	}
}
