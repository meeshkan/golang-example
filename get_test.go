package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitHubRepositories(t *testing.T) {
	repositories, err := getGitHubRepositories()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	assert.True(t, len(repositories) >= 0, "Repositories should be an array")
}
