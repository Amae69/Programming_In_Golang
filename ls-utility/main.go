// Using golang, write an ls -> Linux utility to print out the content of a directory

// 1. ls will take an arguments files or directories
// 2. ls will list the contents of the current directory
// 3. Then ls will prints it to the screen

package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Takes argument
	args := os.Args
	f := args[1]

	//Read the named dir passed on the command line argument and check for error
	c, err := os.ReadDir(f)
	check(err)

	//Print the content on the screen
	fmt.Println(c)

}
