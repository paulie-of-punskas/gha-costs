package main

import (
	"testing"
)

func TestGetPAT(t *testing.T) {
	t.Run("Should return a non empty string, if .env is available.", func(t *testing.T) {
		var got = GetPAT()
		if got == "" {
			t.Errorf("Expected empty string, got %v", got)
		}
	})
}

func TestIsConnectionToGitHubPossible(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		var got = isConnectionToGitHubPossible("paulie-of-punskas", "gha-costs", "")
		var want = true
		if got != want {
			t.Errorf("Expected true to be returned")
		}
	})
}
