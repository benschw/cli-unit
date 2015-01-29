package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var _ = log.Print
var _ = fmt.Print

const (
	BlockPrefix   = "###"
	BlockTypeTest = "### test:"
	BlockTypeWhen = "#### when:"
	BlockTypeThen = "#### then:"
)

type Test struct {
	Title  string
	Script string
	Output string
}

type App struct {
	Tests []Test
}

func (a *App) Parse(files []string) error {
	lines := make([]string, 0)

	for _, filePath := range files {
		l, err := parseFile(filePath)
		if err != nil {
			return err
		}
		lines = append(lines, l...)

	}

	tests, err := parseLines(lines)
	if err == nil {
		a.Tests = tests
	}
	return err
}

func (a *App) Run() error {
	cmd := exec.Command("dig", "any", "google.com")
	out, err := cmd.Output()

	return nil
}

func (a *App) GetDisplay(verbose bool, asJson bool) (string, error) {
	return "", nil
}

func parseFile(filePath string) ([]string, error) {
	lines := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func parseLines(lines []string) ([]Test, error) {
	tests := make([]Test, 0)

	testLines := make([][]string, 0)
	testIdx := -1

	for _, line := range lines {
		if len(line) >= len(BlockTypeTest) && line[0:len(BlockTypeTest)] == BlockTypeTest {
			testIdx += 1
			testLines = append(testLines, make([]string, 0))
		}
		if testIdx >= 0 {
			testLines[testIdx] = append(testLines[testIdx], line)
		}
	}

	for _, lines := range testLines {
		ts, err := parseTestLines(lines)
		if err != nil {
			return tests, err
		}
		tests = append(tests, ts)
	}

	return tests, nil
}

func parseTestLines(lines []string) (Test, error) {
	test := Test{}

	title, err := getTitle(lines)
	if err != nil {
		return test, err
	}
	script, err := getScript(lines)
	if err != nil {
		return test, err
	}
	output, err := getOutput(lines)
	if err != nil {
		return test, err
	}
	test.Title = title
	test.Script = script
	test.Output = output
	return test, nil
}

func getTitle(lines []string) (string, error) {
	if len(lines[0]) < len(BlockTypeTest) {
		return "", errors.New("Test Block Invalid")
	}
	title := lines[0][len(BlockTypeTest):]
	return strings.TrimSpace(title), nil
}

func getScript(lines []string) (string, error) {
	scriptLines := make([]string, 0)
	capture := false
	for _, line := range lines {
		if capture {
			if len(line) >= len(BlockPrefix) && line[0:len(BlockPrefix)] == BlockPrefix {
				return strings.Join(scriptLines, "\n"), nil
			}
			if len(line) >= 1 && line[0:1] == "\t" {
				scriptLines = append(scriptLines, line[1:])
			}
		}
		if len(line) >= len(BlockTypeWhen) && line[0:len(BlockTypeWhen)] == BlockTypeWhen {
			capture = true
		}
	}
	return "", errors.New("Script Block Not Found")
}

func getOutput(lines []string) (string, error) {
	outputLines := make([]string, 0)
	capture := false
	for _, line := range lines {
		if capture {

			if len(line) >= len(BlockPrefix) && line[0:len(BlockPrefix)] == BlockPrefix {
				return strings.Join(outputLines, "\n"), nil
			}
			if len(line) >= 1 && line[0:1] == "\t" {
				outputLines = append(outputLines, line[1:])
			}
		}
		if len(line) >= len(BlockTypeThen) && line[0:len(BlockTypeThen)] == BlockTypeThen {
			capture = true
		}
	}
	if len(outputLines) > 0 {
		return strings.Join(outputLines, "\n"), nil
	}
	return "", errors.New("Output Block Not Found")
}
