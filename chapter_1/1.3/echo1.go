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
	strings.Join(os.Args[1:], " ")
	fmt.Println(time.Since(start))
}
