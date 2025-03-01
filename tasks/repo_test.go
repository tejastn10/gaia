package tasks_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/tejastn10/gaia/tasks"
)

func setupGitRepo(t *testing.T) (string, func()) {
	t.Helper()

	// Create a temporary directory for the Git repo
	tempDir, err := os.MkdirTemp("", "test-repo-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	// Initialize Git repo
	cmd := exec.Command("git", "init", tempDir)
	err = cmd.Run()
	if err != nil {
		t.Fatalf("failed to initialize git repo: %v", err)
	}

	// Return cleanup function to remove the temp dir
	return tempDir, func() {
		os.RemoveAll(tempDir)
	}
}

func TestIsGitRepo(t *testing.T) {
	t.Run("inside a git repo", func(t *testing.T) {
		_, cleanup := setupGitRepo(t)
		defer cleanup()

		err := tasks.IsGitRepo()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("outside a git repo", func(t *testing.T) {
		// Create a temporary directory outside the test repo
		tmpDir := t.TempDir()

		// Change to the new directory to simulate being outside a Git repo
		originalDir, _ := os.Getwd()
		// Restore original directory after test
		defer func() {
			err := os.Chdir(originalDir)
			if err != nil {
				t.Logf("failed to restore original directory: %v", err)
			}
		}()
		err := os.Chdir(tmpDir)
		if err != nil {
			t.Fatalf("failed to change to temp directory: %v", err)
		}

		err = tasks.IsGitRepo()
		if err == nil {
			t.Errorf("expected error, got none")
		}
	})
}

func TestIsGitClean(t *testing.T) {
	t.Run("clean git repo", func(t *testing.T) {
		_, cleanup := setupGitRepo(t)
		defer cleanup()

		clean, err := tasks.IsGitClean()
		if err != nil || !clean {
			t.Errorf("expected clean repo, got err: %v, clean: %v", err, clean)
		}
	})
}
