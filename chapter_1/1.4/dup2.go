// This program finds duplicated lines. It prints the number of times the line
// appears, and which files the lines are present in
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type lineMetaData struct {
	count int
	files map[string]bool
}

func main() {
	counts := make(map[string]lineMetaData)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "STDIN")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, meta := range counts {
		if meta.count > 1 {
			var filenames []string
			for filename, _ := range meta.files {
				filenames = append(filenames, filename)
			}
			fmt.Printf("%d\t%s, found in %s\n", meta.count, line, strings.Join(filenames, ", "))
		}
	}
}
func countLines(f *os.File, counts map[string]lineMetaData, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lineMeta := counts[input.Text()]
		if lineMeta.files == nil {
			lineMeta.files = make(map[string]bool)
		}
		lineMeta.count++
		lineMeta.files[filename] = true
		counts[input.Text()] = lineMeta
	}
}
