// Binary Print prints the binary output of integers, counting upward
// $ ./binary_print 258
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	q := []string{"1"}
	var b, left, right string
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error converting input argument to integer")
	}
	for i := 0; i < max; i++ {
		b = q[0]
		padding := int(math.Ceil(float64(len(b)) / 8) * 8)
		fmt.Printf("%0*s\n", padding, b)
		q = q[1:]
		left = b + "0"
		right = b + "1"
		q = append(q, left)
		q = append(q, right)
	}
}

