package main

import "testing"

func TestIsNiceStringTwo(t *testing.T) {
	t.Run("string is nice", func(t *testing.T) {
		input := "qjhvhtzxzqqjkmpb"

		want := true
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})

	t.Run("string is bad", func(t *testing.T) {
		input := "aaa"

		want := false
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})

	t.Run("string is nice", func(t *testing.T) {
		input := "xxyxx"

		want := true
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})

	t.Run("string is bad", func(t *testing.T) {
		input := "uurcxstgmygtbstg"

		want := false
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})

	t.Run("string is bad", func(t *testing.T) {
		input := "ieodomkazucvgmuy"

		want := false
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})

	t.Run("string is nice", func(t *testing.T) {
		input := "aaaa"

		want := true
		got := IsNiceStringVersionTwo(input)
		assertString(t, want, got, input)
	})
}

func assertString(t testing.TB, want, got bool, input string) {
	if want != got {
		t.Errorf("wanted %v, got %v for %v", want, got, input)
	}
}
