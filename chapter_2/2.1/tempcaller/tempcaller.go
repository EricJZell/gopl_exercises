// tempcaller.go uses the tempconv package to convert temperatures
// This uses the tempconv package in the relative directory because I ran
// go mod edit -replace go.ericjzell.com/tempconv=../tempconv
package main

import (
	"fmt"

	"go.ericjzell.com/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("212°F is %v\n", tempconv.FToK(212))
	fmt.Printf("0k is %v\n", tempconv.KToF(0))
}
