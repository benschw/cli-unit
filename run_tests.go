package main

import (
	"fmt"
	"log"
)

var _ = log.Print
var _ = fmt.Print

func RunTests(tests chan Test, errors chan error, verbose bool) {

	passed := 0
	failed := 0

	for test := range tests {
		if test.Exit {
			sometimesPlural := "test"
			if passed+failed > 1 {
				sometimesPlural = "tests"
			}
			if failed == 0 {
				fmt.Printf("Pass (%d/%d %s successful)\n", passed, passed+failed, sometimesPlural)
			} else {
				fmt.Printf("Fail (%d/%d %s successful)\n", passed, passed+failed, sometimesPlural)
			}

			errors <- nil
			return
		}
		if err := test.Run(); err != nil {
			errors <- err
			return
		}

		testOk := test.Pass()
		if testOk {
			passed += 1
		} else {
			failed += 1
		}

		if verbose || !testOk {
			if testOk {
				fmt.Printf("--- OK: %s\n", test.Title)
			} else {
				fmt.Printf("--- FAIL: %s\n%s\n", test.Title, test.GetFailMessage())
			}
		}
	}
}
