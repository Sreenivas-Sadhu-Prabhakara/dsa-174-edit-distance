package main

import "testing"

func TestEditDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "abc", 3},
		{"abc", "", 3},
		{"", "", 0},
		{"same", "same", 0},
		{"a", "b", 1},
		{"kitten", "sitting", 3},
	}
	for _, c := range cases {
		if got := EditDistance(c.a, c.b); got != c.want {
			t.Errorf("EditDistance(%q,%q) = %d, want %d", c.a, c.b, got, c.want)
		}
		if got := EditDistance(c.b, c.a); got != c.want {
			t.Errorf("EditDistance must be symmetric for %q,%q", c.a, c.b)
		}
	}
}
