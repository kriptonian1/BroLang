package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/kriptonian1/BroLang/src/repl"
)

func init() {

	version := "0.0.1" // Define the version of BroLang

	// Define the flags for the CLI
	showVersion := flag.Bool("v", false, "Displays the version of BroLang you are using")
	showHelp := flag.Bool("h", false, "Displays the help message")

	aliasVersion := flag.Bool("version", false, "Alias for -v")
	aliasHelp := flag.Bool("help", false, "Alias for -h")

	flag.Parse() // Parse the flags

	// Check if the flags are true
	if *showVersion || *aliasVersion {
		fmt.Printf("You are using BroLang v%s\n", version)
		os.Exit(0)
	}

	if *showHelp || *aliasHelp {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s👋 This is the BroLang programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
