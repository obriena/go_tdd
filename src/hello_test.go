package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "EN")
		want := "Hello, Chris"

		assertCorectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		testName := ""
		got := Hello(testName, "EN")
		want := "Hello, World"

		assertCorectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		testName := "Jose"
		got := Hello(testName, "ES")
		want := "Hola, Jose"

		assertCorectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		testName := "Francois"
		got := Hello(testName, "FR")
		want := "Bonjour, Francois"

		assertCorectMessage(t, got, want)
	})
	t.Run("in chinese", func(t *testing.T) {
		testName := "小鱼"
		got := Hello(testName, "CH")
		want := "你好，小鱼"

		assertCorectMessage(t, got, want)
	})
}

func assertCorectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
