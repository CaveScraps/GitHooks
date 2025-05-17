package githooks

import (
	"fmt"
	"testing"
	"os/exec"
)

const (
	testFolder = "test_folder"
)

var (
	testFiles  = []string{"test_file.txt", "test_file2.txt"}
)


func TestGetChangedFiles(t *testing.T){
	// Stash any current changes to avoid conflicts
	exec.Command("git", "stash").Run()
	// Create a folder and a file to test the function
	exec.Command("mkdir", "test_folder").Run()
	exec.Command("touch", "test_folder/test_file.txt").Run()
	exec.Command("cd", "test_folder").Run()
	// Start a git repository
	exec.Command("git", "init").Run()
	exec.Command("git", "add", "test_folder/test_file.txt").Run()
	exec.Command("git", "commit", "-m", "test commit").Run()

	// Call the function to get changed files
	fmt.Println("Files changed on main:")
	files, err := GetChangedFiles("main")
	if err != nil{
		fmt.Println(err)
	}

	for _, element := range files{
		fmt.Println("\t" + element)
	}
	// Clean up
	exec.Command("cd", "-").Run()
	exec.Command("rm", "-rf", "test_folder").Run()
	exec.Command("git", "stash", "pop").Run()
}
