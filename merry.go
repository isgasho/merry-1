package main

import (
	. "github.com/phodal/merry/cmd"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := NewRootCmd(output)
	_ = rootCmd.Execute()
}
