package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var _ = log.Print
var _ = fmt.Print

func RunTests(tests chan Test, errors chan error, verbose bool) {

	passed := 0
	failed := 0

	for test := range tests {
		if test.Exit {
			if failed == 0 {
				fmt.Printf("Pass (%d/%d tests successful)\n", passed, passed+failed)
			} else {
				fmt.Printf("Fail (%d/%d tests successful)\n", passed, passed+failed)
			}

			errors <- nil
			return
		}
		if err := runTest(&test); err != nil {
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
				fmt.Printf("--- FAIL: %s\n%s\n", test.Title, test.Diff())
			}
		}
	}
}

func runTest(test *Test) error {
	bash := exec.Command("bash")
	bash.Stdin = strings.NewReader(test.Script)

	var out bytes.Buffer
	bash.Stdout = &out

	err := bash.Run()
	if err != nil {
		return err
	}

	test.FoundOutput = out.String()
	return nil
}
