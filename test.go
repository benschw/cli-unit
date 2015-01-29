package main

import (
	"github.com/kr/pretty"
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

func (t *Test) Diff() string {
	diff := pretty.Diff(t.ExpectedOutput, t.FoundOutput)
	return strings.Join(diff, "\n")
}

func (t *Test) Pass() bool {

	return strings.TrimSpace(t.ExpectedOutput) == strings.TrimSpace(t.FoundOutput)
}
