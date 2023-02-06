package main

import "testing"

func TestHelloWorld(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", "Hello, World!"},
	}
	for _, c := range cases {
		got := getStringToPrint()
		if got != c.want {
			t.Errorf("got string %s, want %s", got, c.want)
		}
	}
}
