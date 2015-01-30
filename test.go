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

func (t *Test) Diff() string {
	diff := pretty.Diff(t.ExpectedOutput, t.FoundOutput)
	return strings.Join(diff, "\n")
}

func (t *Test) Pass() bool {

	return strings.TrimSpace(t.ExpectedOutput) == strings.TrimSpace(t.FoundOutput)
}
