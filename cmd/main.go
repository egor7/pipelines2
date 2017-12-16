// Samples from https://blog.golang.org/pipelines
// https://golang.org/doc/articles/race_detector.html

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pipelines"
)

var p1 = flag.Int("p1", 1000, "parameter p1")

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "  pipelines [flags] arg0 \n")
	fmt.Fprintf(os.Stderr, "flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if flag.NArg() != 1 {
		usage()
	}
	log.Println(args[0])

	pipelines.gen(1, 2, 3)
}
