// This program echos the command line arguments. It prints each argument
// on a new line
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Running %s\n", os.Args[0])
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", index + 1, arg)
	}
}
