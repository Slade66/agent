package main

import "testing"

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			name: "same base ABC",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "same base AB",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			name: "no common divisor string 1",
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
		{
			name: "no common divisor string 2",
			str1: "AAAAAB",
			str2: "AAA",
			want: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := gcdOfStrings(tc.str1, tc.str2)
			if got != tc.want {
				t.Errorf("gcdOfStrings(%q, %q) = %q, want %q", tc.str1, tc.str2, got, tc.want)
			}
		})
	}
}
