package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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

func ParseFiles(files []string, testConfigs chan Test, errors chan error) {
	lines := make([]string, 0)

	for _, filePath := range files {
		ls, err := parseFile(filePath)
		if err != nil {
			errors <- err
			return
		}
		lines = append(lines, ls...)
	}

	err := parseLines(lines, testConfigs)
	if err != nil {
		errors <- err
	}
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

func parseLines(lines []string, tests chan Test) error {

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
		test, err := parseTestLines(lines)
		if err != nil {
			return err
		}

		tests <- test
	}
	tests <- Test{Exit: true}
	return nil
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
	test.ExpectedOutput = output
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
