package tasks

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestRepo(t *testing.T) func() {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "git-test-repo")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	cmds := [][]string{
		{"git", "init"},
		{"git", "config", "user.name", "Test User"},
		{"git", "config", "user.email", "test@example.com"},
		{"sh", "-c", "echo 'initial' > file.txt"},
		{"git", "add", "file.txt"},
		{"git", "commit", "-m", "Initial commit"},
	}

	for _, cmd := range cmds {
		c := exec.Command(cmd[0], cmd[1:]...)
		c.Dir = tempDir
		if err := c.Run(); err != nil {
			t.Fatalf("Failed to execute %v: %v", cmd, err)
		}
	}

	prevDir, _ := os.Getwd()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	return func() {
		os.Chdir(prevDir)
		os.RemoveAll(tempDir)
	}
}

func TestGetFirstCommitHash(t *testing.T) {
	cleanup := setupTestRepo(t)
	defer cleanup()

	ctx := context.Background()
	commitHash, err := GetFirstCommitHash(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, commitHash)
}

func TestUpdateFirstCommit(t *testing.T) {
	cleanup := setupTestRepo(t)
	defer cleanup()

	ctx := context.Background()

	// Mock user input by replacing os.Stdin with a pipe containing "no\n"
	input := bytes.NewBufferString("no\n")
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	// Restore original stdin after test
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r
	w.Write(input.Bytes())
	w.Close()

	err := UpdateFirstCommit(ctx, "Updated first commit")
	assert.NoError(t, err)
}
