package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var testCases = []struct {
		path          string
		expectedCount int
		expectedErr   error
	}{
		{"testdata/input.txt", 7, nil},
	}
	for _, tc := range testCases {
		count, err := solve(tc.path)
		if err != tc.expectedErr {
			t.Fatal(err)
		}
		if count != tc.expectedCount {
			t.Fatalf("expected %d , got %d instead.\n", tc.expectedCount, count)
		}
	}
}
