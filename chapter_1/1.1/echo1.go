// This program echos the command line arguments, and prints the command that
// invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Running %s\n", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
