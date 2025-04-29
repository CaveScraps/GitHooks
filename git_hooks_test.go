package githooks

import (
	"fmt"
	"testing"
)

func TestPrintGetChangedFiles(t *testing.T){
	fmt.Println("Files changed on main:")
	files, err := GetChangedFiles("main")
	if err != nil{
		fmt.Println(err)
	}

	for _, element := range files{
		fmt.Println("\t" + element)
	}
}
