package main

import "testing"

func TestValidUtf8(t *testing.T) {
	cases := []struct {
		data     []int
		expected bool
	}{
		{[]int{235, 140, 4}, false},
		{[]int{197, 130, 1}, true},
	}
	for _, tc := range cases {
		actual := validUtf8(tc.data)
		if actual != tc.expected {
			t.Errorf("expected %v for %0xv, got %v", tc.expected, tc.data, actual)
		}
	}
}
