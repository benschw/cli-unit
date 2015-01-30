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
	BlockTypeGenericHeader = "#"
	BlockPrefix            = "###"
	BlockTypeTest          = "### test:"
	BlockTypeWhen          = "#### when:"
	BlockTypeThen          = "#### then:"
	StrictFlag             = "# strict"
)

func ParseFiles(files []string, tests chan Test, errors chan error) {

	for _, filePath := range files {

		lines, err := fileToLines(filePath)
		if err != nil {
			errors <- err
		}

		if err = generateTestsFromLines(lines, tests); err != nil {
			errors <- err
		}
	}

	tests <- Test{Exit: true}
}

func fileToLines(filePath string) ([]string, error) {
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
	var buffer []string
	capture := false

	for _, line := range lines {
		if lineIsGenericHeader(line) || lineIsTestBlockHeader(line) {
			if capture {
				test, err := getTest(buffer)
				if err != nil {
					return err
				}
				tests <- test
			}
		}
		if lineIsTestBlockHeader(line) {
			capture = true
			buffer = make([]string, 0)
		}
		if lineIsGenericHeader(line) {
			capture = false
		}
		if capture {
			buffer = append(buffer, line)
		}
	}

	return nil
}

func getTest(lines []string) (Test, error) {
	test := Test{}

	title, err := getTitle(lines)
	if err != nil {
		return test, err
	}
	test.Title = title

	whenBlock, err := getBlock(lines, BlockTypeWhen)
	if err != nil {
		return test, err
	}
	test.Script = strings.Join(whenBlock, "\n")

	thenBlock, err := getBlock(lines, BlockTypeThen)
	if err != nil {
		return test, err
	}
	test.ExpectedOutput = strings.Join(thenBlock, "\n")

	isStrict, err := isStrict(lines)
	if err != nil {
		return test, err
	}

	test.Strict = isStrict
	return test, nil
}

func getTitle(lines []string) (string, error) {
	if len(lines[0]) < len(BlockTypeTest) {
		return "", errors.New("Test Block Invalid")
	}
	title := lines[0][len(BlockTypeTest):]
	return strings.TrimSpace(title), nil
}
func isStrict(lines []string) (bool, error) {
	thenBlock, err := getBlock(lines, BlockTypeThen)
	if err != nil {
		return false, err
	}
	trimmed := strings.Split(strings.TrimSpace(strings.Join(thenBlock, "\n")), "\n")
	return trimmed[len(trimmed)-1] == StrictFlag, nil
}
func getBlock(lines []string, blockType string) ([]string, error) {
	buffer := make([]string, 0)
	capture := false
	for _, line := range lines {
		if capture {
			if lineIsBlockHeader(line) {
				return buffer, nil
			}
			if lineIsBlockBody(line) {
				buffer = append(buffer, line[1:])
			}
		}
		if lineIsHeader(line, blockType) {
			capture = true
		}
	}
	if len(buffer) > 0 {
		return buffer, nil
	}
	return buffer, errors.New("Bad format: expected block missing")
}

func lineIsBlockHeader(line string) bool {
	return lineIsHeader(line, BlockPrefix)
}

func lineIsBlockBody(line string) bool {
	return lineIsHeader(line, "\t")
}

func lineIsTestBlockHeader(line string) bool {
	return lineIsHeader(line, BlockTypeTest)
}

func lineIsWhenBlockHeader(line string) bool {
	return lineIsHeader(line, BlockTypeWhen)
}

func lineIsThenBlockHeader(line string) bool {
	return lineIsHeader(line, BlockTypeThen)
}

func lineIsHeader(line string, headerType string) bool {
	return len(line) >= len(headerType) && line[0:len(headerType)] == headerType
}

func lineIsGenericHeader(line string) bool {
	if lineIsHeader(line, BlockTypeGenericHeader) {
		if !lineIsTestBlockHeader(line) && !lineIsWhenBlockHeader(line) && !lineIsThenBlockHeader(line) {
			return true
		}
	}
	return false
}
