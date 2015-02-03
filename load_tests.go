package main

import (
	"fmt"
	"log"
)

var _ = log.Print
var _ = fmt.Print

func LoadTests(files []string, tests chan Test, errors chan error) {

	for _, filePath := range files {
		log.Println(filePath)
		parser, err := NewTestFileParser(filePath)
		if err != nil {
			errors <- err
		}

		test := &Test{}
		for test != nil {
			test, err = parser.NextTest()
			if err != nil {
				errors <- err
				return
			}
			if test != nil {
				tests <- *test
			}
		}

	}
	tests <- Test{Exit: true}

}
