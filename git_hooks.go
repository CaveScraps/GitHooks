package githooks

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
* GetChangedFiles retrieves the list of files changed in the current branch compared to the target branch.
* Note these are only files that have been marked as changed by git, i.e. Committed files.
* @param targetBranchName: The name of the target branch to compare against.
* @return: A slice of strings containing the names of the changed files.
* @error: An error if the git command fails to execute.
 */
func GetChangedFiles(targetBranchName string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", targetBranchName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	splitStrings := strings.Split(string(output), "\n")
	return splitStrings, nil
}
