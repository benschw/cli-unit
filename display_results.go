package main

import (
	"fmt"
	"log"
)

var _ = log.Print
var _ = fmt.Print

func DisplayResults(tests chan Test, errors chan error, verbose bool, asJson bool) {
	// log.Printf("%+v", tests)
	for test := range tests {
		if test.Exit {
			errors <- nil
			return
		}
		fmt.Println(test.Diff())

	}
}
