package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		days   int
		expOut int64
		expErr error
		solve  func(string, int) (int64, error)
	}{
		{"day06/part1", "testdata/input.txt", 80, 5934, nil, solve},
		{"day06/part2", "testdata/input.txt", 256, 26984457539, nil, solve},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := tc.solve(tc.input, tc.days)
			if err != tc.expErr {
				t.Fatalf("expected %s , got %s", tc.expErr, err)
			}
			if v != tc.expOut {
				t.Fatalf("expected %d , got %d", tc.expOut, v)
			}
		})
	}
}
