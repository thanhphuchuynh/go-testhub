package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// driver tests
	testcase := []struct {
		x, y int
		res  int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
		{5, 6, 11},
		{6, 7, 13},
	}

	for _, tc := range testcase {
		if res := Add(tc.x, tc.y); res != tc.res {

			t.Errorf("Add(%d, %d) = %d, want %d", tc.x, tc.y, res, tc.res)
		}
	}

}
