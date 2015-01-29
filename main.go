package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var _ = log.Print
var _ = fmt.Print

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	flag.Usage = func() {
		fmt.Printf("Usage: cli-unit [options] test-file [test-file2...]\n\nOptions:\n")
		flag.PrintDefaults()
	}

	verbose := flag.Bool("v", true, "verbose")

	flag.Parse()

	// test files
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	files := make([]string, 0)
	for i := 0; i < flag.NArg(); i++ {
		files = append(files, flag.Arg(i))
	}

	// run app
	tests := make(chan Test)
	errors := make(chan error)

	go ParseFiles(files, tests, errors)
	go RunTests(tests, errors, *verbose)

	err := <-errors

	if err != nil {
		if *verbose {
			log.Println(err)
		}
		os.Exit(1)
	}
}
