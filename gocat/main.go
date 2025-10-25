// Writing a simple cat utility
//NOTE: a basic cat utility will

// 1. Takes file name as argument
// 2. Read from a file
// 3. Print the file content to screen

package main

import (
	"fmt"
	"os"
)

func main() {
	// Takes file name as argument
	args := os.Args

	// Read from a file
	fc, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("An error occured while reading file: %e", err)
	}

	// Print the file content to screen
	fmt.Println(string(fc))

}
