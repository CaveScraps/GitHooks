# GitHooks

GitHooks is a Golang module that provides useful functions to write git hooks.

## Usage

Add the module as a dependency to your golang project:
```bash
go get github.com/CaveScraps/GitHooks
```

Then in your golang githook:
```go
package main

import (
	"github.com/CaveScraps/GitHooks"
	"fmt"
)

func main() {
	fmt.Println(githooks.GetChangedFiles("main"))
}
```

## Available Functions
### GetChangedFiles
Given a branch name as a target, this function will return all the (committed) files changed on the current branch

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[WTFPL](https://choosealicense.com/licenses/wtfpl/)
