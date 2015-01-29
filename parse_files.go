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

	lines, err := getFileContentsAsLines(files)
	if err != nil {
		errors <- err
	}

	if err = generateTestsFromLines(lines, testConfigs); err != nil {
		errors <- err
	}
}

func getFileContentsAsLines(files []string) ([]string, error) {
	lines := make([]string, 0)
	for _, filePath := range files {
		ls, err := parseFile(filePath)
		if err != nil {
			return lines, err
		}
		lines = append(lines, ls...)
	}
	return lines, nil
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

func generateTestsFromLines(lines []string, tests chan Test) error {

	testLines := make([][]string, 0)
	idx := -1
	for _, line := range lines {
		if lineIsTestBlockHeader(line) {
			idx += 1
			testLines = append(testLines, make([]string, 0))
		}
		testLines[idx] = append(testLines[idx], line)
	}

	for _, ls := range testLines {
		test, err := generateTestFromLines(ls)
		if err != nil {
			return err
		}
		tests <- test
	}

	tests <- Test{Exit: true}
	return nil
}

func generateTestFromLines(lines []string) (Test, error) {
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
			if lineIsBlockHeader(line) {
				return strings.Join(scriptLines, "\n"), nil
			}
			if lineIsBlockBody(line) {
				scriptLines = append(scriptLines, line[1:])
			}
		}
		if lineIsWhenBlockHeader(line) {
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
			if lineIsBlockHeader(line) {
				return strings.Join(outputLines, "\n"), nil
			}
			if lineIsBlockBody(line) {
				outputLines = append(outputLines, line[1:])
			}
		}
		if lineIsThenBlockHeader(line) {
			capture = true
		}
	}
	if len(outputLines) > 0 {
		return strings.Join(outputLines, "\n"), nil
	}
	return "", errors.New("Output Block Not Found")
}

func lineIsBlockHeader(line string) bool {
	return len(line) >= len(BlockPrefix) && line[0:len(BlockPrefix)] == BlockPrefix
}

func lineIsBlockBody(line string) bool {
	return len(line) >= 1 && line[0:1] == "\t"
}

func lineIsTestBlockHeader(line string) bool {
	return len(line) >= len(BlockTypeTest) && line[0:len(BlockTypeTest)] == BlockTypeTest
}

func lineIsWhenBlockHeader(line string) bool {
	return len(line) >= len(BlockTypeWhen) && line[0:len(BlockTypeWhen)] == BlockTypeWhen
}

func lineIsThenBlockHeader(line string) bool {
	return len(line) >= len(BlockTypeThen) && line[0:len(BlockTypeThen)] == BlockTypeThen
}
