package commando

import (
	"fmt"
	"errors"
	"os"
	"text/tabwriter"
	"strings"
	"regexp"
)

var tw *tabwriter.Writer

// Command is the base type for all commands.
type Command struct {
	Name string                  // Name of command, typically how a command is called from the cli.
	Description string           // A Description of the command, printed in usage.
	Options map[string]*Option   // A map of the flags attached to this command, they are looked up by their name.
	Children map[string]*Command // A map of all the subcommands, looked up by their name.
	Parent *Command              // A pointer to the command's parent.  not set in root command.
	Execute func()               // The function to run when executing a command.

}

// Option is the type for flag options like "-p" or "--path"
type Option struct {
	Name string        // Name of Option, its name is used to retrieve its value.
	Description string // A Description of the option, used when printing usage.
	Flags []string     // The flags associated with the option.
	Value interface{}  // Where the value of a given flag is scanned into.
	Present bool       // Used to determine whether or not a flag is present, typically for a bool type flag.
	Required bool      // If a flag is required and not present, usage for owning command is printed.
}

var tabWriter *tabwriter.Writer

// AddSubcommand attaches a command to a parent, as well as sets the parent property on the child command.
// Commands can be limitlessly nested (though, I don't recommend it).
func (c *Command) AddSubCommand(child *Command) {
	if c.Children == nil {
		c.Children = make(map[string]*Command)
	}
	child.Parent = c
	c.Children[child.Name] = child
}

// PrintHelp is used to print info and usage for any command.
// It knows if a command is the last in the chain, and if so, prints usage with just Options (Flags)
func (c *Command) PrintHelp() {
	if c.hasChildren() {
		fmt.Println("\nUsage:", c.Name, "COMMAND [args..]\n")
		fmt.Println(c.Description, "\n")
		fmt.Println("Commands:")
		for _, cmd := range(c.Children) {
			PrintFields(true, 4, cmd.Name, cmd.Description)
		}
	} else {
		fmt.Printf("\nUsage: %s %s [options...]\n\n", c.Parent.Name, c.Name)
		fmt.Println(c.Description)
		fmt.Println("\nOptions:")
		for _, opt := range c.Options {
			PrintFields(true, 4, strings.Join(opt.Flags, ", "), opt.Description)
		}
	
	}
}

// hasChildren is a private method that determines whether or not a command has children.
// Parse uses hasChildren to decide whether or not to continue recursing.
func (c *Command) hasChildren() bool {
	if c.Children != nil {
		return true
	} else {
		return false
	}
}

// AddOption is used to add an option (Flag) to a command.
// Ex: cmd.AddOption("path", "Path to a thing", true, "-p", "--path")
func (c *Command) AddOption(name string, descrip string, req bool, flags ...string) {
	if c.Options == nil {
		c.Options = make(map[string]*Option)
	}
	opt := &Option{
		Name:        name,
		Description: descrip,
		Flags:       flags,
		Required:    req,
	}
	c.Options[name] = opt
}

// PrintFields is a wrapper for an IO Writer / Formatter.
// Using commando.PrintFields evenly spaces output into columns.
func PrintFields(indent bool, width int, fields ...interface{}) {
	argArray := make([]interface{}, 0)
	if indent {
		argArray = append(argArray, strings.Repeat(" ", width))
	}
	for i, field := range fields {
		argArray = append(argArray, field)
		if i < (len(fields) - 1) {
			argArray = append(argArray, "\t")
		}
	}
	fmt.Fprintln(tw, argArray...)
}

// Parse is the entry point into Commando.
// It recurses all the children of a command, finally executing the last command in the chain.
func (c *Command) Parse() {

	tw = tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	defer tw.Flush()

	if len(os.Args) == 1 {
		c.PrintHelp()
		return
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		c.PrintHelp()
		return
	}

	if err := c.setOptions(); err == nil {
		c.executeChildren()
	} else {
		c.PrintHelp()
	}
}

// findChild is a private method used to locate the requested child command of a parent.
// It is used in the recursive lookup in Parse
func (c *Command) findChild() *Command {
	var child *Command
	for _, arg := range os.Args {
		if c.Children[arg] != nil {
			child = c.Children[arg]
		}
	}
	return child
}

// setOptions is used to retrieve flagged options  and set their values.
func (c *Command) setOptions() error {
	for i, arg := range os.Args {
		for _, opt := range c.Options {
			for _, flag := range opt.Flags {
				if match, _ := regexp.MatchString(arg, flag); match {
					opt.Value = os.Args[i+1]
					opt.Present = true
				}
			}
		}
	}
	for _, opt := range c.Options {
		if opt.Required && opt.Value == nil {
			err := errors.New("required option missing")
			return err
		}
	}
	return nil
}

// executeChildren is the recurive part of Parse.
// It determines if a command has children, and if it does, executes them.
// If not, it continues to recurse.
func (c *Command) executeChildren() {
	r, _ := regexp.Compile("^-{1,2}.*")
	if !r.MatchString(os.Args[1]) {
		if c.hasChildren() {
			if child := c.findChild(); child != nil {
				child.Parse()
			} else {
				c.PrintHelp()
			}
			return
		} else {
			c.Execute()
		}
	}
}
