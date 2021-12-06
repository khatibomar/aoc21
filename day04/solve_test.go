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
		{"Day04/part1/input1", "testdata/input1.txt", 4512, nil, solve},
		{"Day04/part1/input2", "testdata/input2.txt", 27027, nil, solve},
		{"Day04/part2/input1", "testdata/input1.txt", 1924, nil, solve2},
		{"Day04/part2/input2", "testdata/input2.txt", 36975, nil, solve2},
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
