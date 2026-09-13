package main

import (
	"testing"
)

func TestGetPAT(t *testing.T) {
	t.Run("Should return a non empty string, if .env is available.", func(t *testing.T) {
		var got = GetPAT()
		if got == "" {
			t.Errorf("Expected non empty string, got: '%v'", got)
		}
	})
}

func TestIsConnectionToGitHubPossible(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		var got = IsConnectionToGitHubPossible("paulie-of-punskas/gha-costs", false)
		var want = true
		if got != want {
			t.Errorf("Expected true to be returned")
		}
	})
}
