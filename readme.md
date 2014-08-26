Commando
========

A CLI Parser for Go

[![Build Status](https://travis-ci.org/danielscottt/commando.svg?branch=master)](https://travis-ci.org/danielscottt/commando)

godocs: http://godoc.org/github.com/danielscottt/commando

## Overview

Commando is a cli parser that handles nested commands, usage / help output, flags, and output _formatting_ for you.

Define a root command, and attach subcommands to it.  Tell your new commands what function to execute, and that's it.

## Why

Because I don't like the UX of Flags.  The goal here is a clean API to define complex cli programs.

## Example

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

See the [godocs](http://godoc.org/github.com/danielscottt/commando) for further documentation.
