package commando

import (
	"fmt"
	"os"
	"testing"
)

var root *Command

func setupDummyTree() {
	root = testDummyCommand("root")
	sub := testDummyCommand("sub")
	root.AddSubCommand(sub)
}

func testDummyCommand(name string) *Command {
	cmd := &Command{
		Name:        name,
		Description: "A Stub Command",
	}
	return cmd
}

func TestAddSubCommandChildren(t *testing.T) {

	setupDummyTree()

	if root.Children["sub"] == nil {
		t.Fatalf("Expected: sub to be added to root command's childen map. Got: nil")
	}

}

func TestAddSubCommandParent(t *testing.T) {

	setupDummyTree()

	if root.Children["sub"].Parent == nil {
		t.Fatalf("Expected: root to be added as Parent to sub. Got: nil")
	}
}

func TestAddOption(t *testing.T) {
	root = testDummyCommand("root")
	root.AddOption("path", "A path to a thing", true, "-p", "--path")
	if root.Options["path"] == nil {
		t.Fatalf("Exepected: option \"path\" to be added to root.Options. Got: nil")
	}
}

func testStandaloneExec() {
	fmt.Println("sub ran")
}

func ExampleParse() {
	os.Args = []string{"root", "sub", "--path", "/test/stub"}
	setupDummyTree()
	root.Children["sub"].Execute = testStandaloneExec
	root.Parse()
	// Output:
	// sub ran
}

func testOptionExec() {
	fmt.Println("sub ran with path", root.Children["sub"].Options["path"].Value)
}

func ExampleParseWithOption() {
	os.Args = []string{"root", "sub", "--path", "/test/stub"}
	setupDummyTree()
	root.Children["sub"].Execute = testOptionExec
	root.Children["sub"].AddOption("path", "A path to a thing", true, "--path")
	root.Parse()
	// Output:
	// sub ran with path /test/stub
}
