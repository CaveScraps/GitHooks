package githooks

import (
	"fmt"
	"os/exec"
)

func GetChangedFiles(targetBranchName string) {
	cmd := exec.Command("git", "diff", "--name-only", targetBranchName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}
