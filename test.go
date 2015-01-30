package main

import (
	"bytes"
	"github.com/kr/pretty"
	"os/exec"
	"strings"
)

type Test struct {
	Title          string
	Script         string
	FoundOutput    string
	ExpectedOutput string
	Strict         bool
	Exit           bool
}

func (t *Test) Run() error {
	bash := exec.Command("bash")
	bash.Stdin = strings.NewReader(t.Script)

	var out bytes.Buffer
	bash.Stdout = &out

	err := bash.Run()
	if err != nil {
		return err
	}

	t.FoundOutput = out.String()
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

func (t *Test) Diff() string {
	diff := pretty.Diff(t.GetExpectedOutput(), t.GetFoundOutput())
	return strings.Join(diff, "\n")
}

func (t *Test) Pass() bool {
	return t.GetExpectedOutput() == t.GetFoundOutput()
}
