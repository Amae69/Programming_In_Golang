package main

import (
	"flag"
	"fmt"
	"os"

	"gocat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocat <command> [options] <file>")
		fmt.Println("Commands: cat, head, tail")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "cat":
		catCmd := flag.NewFlagSet("cat", flag.ExitOnError)
		number := catCmd.Bool("n", false, "show line numbers")
		limit := catCmd.Int("l", 0, "limit number of lines to print (0 = all)")
		catCmd.Parse(os.Args[2:])
		output, err := gocat.RunCat(catCmd, *number, *limit)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Print(output)

	case "head":
		headCmd := flag.NewFlagSet("head", flag.ExitOnError)
		lines := headCmd.Int("n", 10, "number of lines to show")
		headCmd.Parse(os.Args[2:])
		output, err := gocat.RunHead(headCmd, *lines)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Print(output)

	case "tail":
		tailCmd := flag.NewFlagSet("tail", flag.ExitOnError)
		lines := tailCmd.Int("n", 10, "number of lines to show")
		tailCmd.Parse(os.Args[2:])
		output, err := gocat.RunTail(tailCmd, *lines)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Print(output)

	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Available commands: cat, head, tail")
		os.Exit(1)
	}
}
