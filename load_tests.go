package main

import (
	"fmt"
	"log"
)

var _ = log.Print
var _ = fmt.Print

func LoadTests(files []string, tests chan Test, errors chan error) {

	for _, filePath := range files {

		parser, err := NewTestFileParser(filePath)
		if err != nil {
			errors <- err
		}

		for {
			test, err := parser.NextTest()
			if err != nil {
				errors <- err
				return
			}
			if test == nil {
				tests <- Test{Exit: true}
				return
			}
			tests <- *test
		}

	}

}
