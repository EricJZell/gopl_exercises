// Fetchall fetches URLs in parallel and reports thier times and sizes.
// This version writes the results to a file
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Create("results.txt")
	if err != nil {
		fmt.Printf("Error opening results file: %v", err)
	}
	for _, url := range os.Args[1:] {
		go fetch (url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		f.WriteString(<-ch) // receive from channel
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	defer f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
