// This program echos the command line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Running %s\n", os.Args[0])

	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Join took %v\n", time.Since(start))

	start = time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s+= sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("For loop took %v\n", time.Since(start))
}
