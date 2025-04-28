package githooks

import (
	"testing"
)

func TestStatus(t *testing.T){
	GetChangedFiles("main");
}
