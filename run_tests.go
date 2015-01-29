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

func RunTests(testConfigs chan Test, runTests chan Test, errors chan error) {
	for test := range testConfigs {
		if test.Exit {
			runTests <- Test{Exit: true}
			return
		}
		if err := runTest(&test); err != nil {
			errors <- err
			return
		}
		if err := calcResults(&test); err != nil {
			errors <- err
			return
		}
		runTests <- test
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

func calcResults(test *Test) error {
	test.Result = test.ExpectedOutput == test.FoundOutput

	return nil
}
