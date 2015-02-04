package main

import (
	"bytes"
	"fmt"
	"github.com/kr/pretty"
	"log"
	"os/exec"
	"strings"
)

var _ = log.Print

type Test struct {
	Title          string
	Script         string
	FoundOutput    string
	FoundError     string
	Err            error
	ExpectedOutput string
	Strict         bool
	Exit           bool
}

func (t *Test) Run() error {
	bash := exec.Command("bash")
	bash.Stdin = strings.NewReader(t.Script)

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	bash.Stdout = &stdOut
	bash.Stderr = &stdErr
	t.Err = bash.Run()

	t.FoundOutput = stdOut.String()
	t.FoundError = stdErr.String()
	return nil
}

func (t *Test) GetExpectedOutput() string {
	if t.Strict {
		return t.ExpectedOutput
	} else {
		return strings.TrimSpace(t.ExpectedOutput)
	}
}
func (t *Test) GetFoundOutput() string {
	if t.Strict {
		return t.FoundOutput
	} else {
		return strings.TrimSpace(t.FoundOutput)
	}
}
func (t *Test) GetFoundError() string {
	return strings.TrimSpace(t.FoundError)
}

func (t *Test) GetFailMessage() string {
	if t.Err != nil {
		return fmt.Sprintf("%s: %s", t.Err, t.GetFoundError())
	} else {
		return t.Diff()
	}
}

func (t *Test) Diff() string {
	diff := pretty.Diff(t.GetExpectedOutput(), t.GetFoundOutput())
	return strings.Join(diff, "\n")
}

func (t *Test) Pass() bool {
	return t.Err == nil && (t.GetExpectedOutput() == t.GetFoundOutput())
}
