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
	Result         bool
	Exit           bool
}

func (t *Test) Diff() string {
	diff := pretty.Diff(t.ExpectedOutput, t.FoundOutput)
	return strings.Join(diff, "\n")
}

func (t *Test) Pass() bool {
	return t.ExpectedOutput == t.FoundOutput
}
