/*
A CLI Parser for Go

Overview

Commando is a cli parser that handles nested commands, usage / help output, flags, and output _formatting_ for you.

Define a root command, and attach subcommands to it.  Tell your new commands what function to execute, and that's it.

Why

Because I don't like the UX of Flags.  The goal here is a clean API to define complex cli programs.
*/
package commando
