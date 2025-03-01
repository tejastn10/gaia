package tasks

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetFirstCommitHash(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "rev-list", "--max-parents=0", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func UpdateFirstCommit(ctx context.Context, message string) error {
	err := IsGitRepo()
	if err != nil {
		return err
	}

	if _, err := GetFirstCommitHash(ctx); err != nil {
		return err
	}

	isGitRepoClean, err := IsGitClean()
	if err != nil {
		return err
	}

	if !isGitRepoClean {
		return errors.New("working tree is not clean, commit or stash changes first")
	}

	// Automate replacing 'pick' with 'edit' in rebase sequence
	rebaseCmd := exec.CommandContext(ctx, "bash", "-c", "GIT_SEQUENCE_EDITOR=\"sed -i.bak 's/^pick /edit /'\" git rebase -i --root")
	err = rebaseCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start interactive rebase: %w", err)
	}

	// Amend commit message
	err = exec.CommandContext(ctx, "git", "commit", "--amend", "-m", message).Run()
	if err != nil {
		if abortErr := exec.CommandContext(ctx, "git", "rebase", "--abort").Run(); abortErr != nil {
			return fmt.Errorf("failed to amend commit: %w (also failed to abort rebase: %v)", err, abortErr)
		}
		return fmt.Errorf("failed to amend commit: %w", err)
	}

	// Continue rebase
	err = exec.CommandContext(ctx, "git", "rebase", "--continue").Run()
	if err != nil {
		if abortErr := exec.CommandContext(ctx, "git", "rebase", "--abort").Run(); abortErr != nil {
			return fmt.Errorf("failed to complete rebase: %w (also failed to abort rebase: %v)", err, abortErr)
		}
		return fmt.Errorf("failed to complete rebase: %w", err)
	}

	fmt.Printf("Successfully updated first commit: \"%s\"\n", message)

	// Ask user for force push
	fmt.Print("Do you want to force push the changes? (yes/no): ")
	var response string
	if _, scanErr := fmt.Scanln(&response); scanErr != nil {
		return fmt.Errorf("failed to read user input: %w", scanErr)
	}

	if strings.ToLower(response) == "yes" {
		err = exec.CommandContext(ctx, "git", "push", "--force").Run()
		if err != nil {
			return fmt.Errorf("failed to force push: %w", err)
		}
		fmt.Println("Changes force pushed successfully.")
	} else {
		fmt.Println("To force push manually, run: git push --force")
	}

	return nil
}
