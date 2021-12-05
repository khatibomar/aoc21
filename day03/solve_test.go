package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expOut int
		expErr error
		solve  func(string) (int, error)
	}{
		{"Day03/part1", "testdata/input.txt", 198, nil, solve},
		{"Day03/part2", "testdata/input.txt", 230, nil, solve2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := tc.solve(tc.input)
			if err != tc.expErr {
				t.Fatalf("expected %s , got %s", tc.expErr, err)
			}
			if v != tc.expOut {
				t.Fatalf("expected %d , got %d", tc.expOut, v)
			}
		})
	}
}
