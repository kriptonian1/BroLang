package repl

import (
	"fmt"
	"os"
)

func cmdHelper(cmd string) {
	switch cmd {
	case ".help":
		fmt.Println("Welcome to the BroLang REPL!")
		fmt.Println("Here are the available commands:")
		fmt.Println(".help - Displays this message")
		fmt.Println(".exit - Exits the REPL")
		fmt.Println(".clear - Clears the screen")
		fmt.Println(".version - Displays the version of BroLang you are using")
		fmt.Println(".license - Displays the license of BroLang")
		fmt.Println(".github - Displays the GitHub repository of BroLang")
		fmt.Println(".website - Displays the website of BroLang")
		fmt.Println(".author - Displays the author of BroLang")
		fmt.Println(".contributors - Displays the contributors of BroLang")
		fmt.Println(".donate - Displays the donation link of BroLang")
	case ".exit":
		os.Exit(0)
	case ".clear":
		fmt.Print("\033[H\033[2J")
	case ".version":
		fmt.Printf("You are using BroLang v%s\n", os.Getenv("VERSION"))
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("Type .help to see the available commands")
	}

}
