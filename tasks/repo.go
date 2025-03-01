package tasks

import (
	"errors"
	"os/exec"
	"strings"
)

func IsGitRepo() error {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	output, err := cmd.Output()
	if err != nil {
		return errors.New("not a git repository")
	}

	if strings.TrimSpace(string(output)) != "true" {
		return errors.New("not inside a git working tree")
	}

	return nil
}

func IsGitClean() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	if len(output) > 0 {
		return false, errors.New("working tree is not clean")
	}

	return true, nil
}
