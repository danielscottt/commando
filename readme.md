Commando
========

A CLI Parser for Go


```go
package main

import (
	"github.com/danielscottt/commando"
)

var root, cmd1 *commando.Command


func runCmd1() {
	commando.PrintFields(false, 0, "HI", "DOOD")
	commando.PrintFields(false, 0, "for path you said:", cmd1.Options["path"].Value)
	commando.PrintFields(false, 0, "jkfdhfkjdkljdshfkjhds", "fdkjkdfhjkdsfkjdsfh")
}

func main() {

	root = &commando.Command{
		Name: "main.go",
		Description: "Testing Commando, a CLI parser",
	}

	cmd1 = &commando.Command{
		Name: "cmd1",
		Description: "Command 1",
		Execute: runCmd1,
	}
	cmd1.AddOption("path", "Path to a thing", true, "-p", "--path")
	root.AddSubCommand(cmd1)

	root.Parse()
}
```
