package main

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	testCases := []struct {
		input  string
		expOut int64
		expErr error
	}{
		{"testdata/input.txt", 198, nil},
	}
	for _, tc := range testCases {
		v, err := solve(tc.input)
		if err != tc.expErr {
			t.Fatalf("expected %s , got %s", tc.expErr, err)
		}
		if v != tc.expOut {
			t.Fatalf("expected %d , got %d", tc.expOut, v)
		}
	}
}
