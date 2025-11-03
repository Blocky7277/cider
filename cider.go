package main

import (
	"fmt"
	"os"
	"cider/cmds"
)

func main() {
	args := os.Args[1:] // Get the arguments without the executable name	
	if len(args) < 1 {
		cmds.Help()
		//ENHANCE: TUI
	} else if args[0] == "help" {
		cmds.Help()
	} else if args[0] == "config" {
		cmds.Config(args[1:])	
	} else if args[0] == "install" {
		cmds.Install()
	} else if args[0] == "uninstall" {
		// cmds.Uninstall()
	} else {
		fmt.Printf("Argument \"%s\" not found \n", args[0])
		cmds.Help()
	}
}
