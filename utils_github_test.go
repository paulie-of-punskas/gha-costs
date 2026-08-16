package main

import (
	"testing"
)

func TestGetPAT(t *testing.T) {
	t.Run("Should return empty string, if .env and GITHUB_PAT are not available.", func(t *testing.T) {
		var got = GetPAT()
		var want = ""
		if got != want {
			t.Errorf("Expected empty string, got %v", got)
		}
	})
}
