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
		{"day01/part1", "testdata/input.txt", 5934, nil, solve},
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
