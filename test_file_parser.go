package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

var _ = log.Print

const (
	BlockTypeGenericHeader = "#"
	BlockPrefix            = "###"
	BlockTypeTest          = "### test:"
	BlockTypeWhen          = "#### when:"
	BlockTypeThen          = "#### then:"
	StrictFlag             = "(strict)"
)

type TestFileParser struct {
	lines []string
	idx   int
}

func NewTestFileParser(filePath string) (*TestFileParser, error) {
	lines := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &TestFileParser{
		lines: lines,
		idx:   0,
	}, nil
}

func (t *TestFileParser) NextTest() (*Test, error) {

	lines := append(t.lines, BlockTypeGenericHeader)
	buffer := make([]string, 0)

	capture := false
	complete := false

	for i := t.idx; i < len(lines) && complete == false; i++ {
		line := lines[i]

		if capture && (lineIsTestBlockHeader(line) || lineIsGenericHeader(line)) {
			complete = true

			test, err := parseTest(buffer)
			if err != nil {
				return nil, err
			}
			t.idx = i - 1
			return &test, nil
		} else {

			if lineIsTestBlockHeader(line) && !capture {
				capture = true
			}

			if capture {
				buffer = append(buffer, line)
			}
		}

	}
	return nil, nil
}

func parseTest(lines []string) (Test, error) {
	test := Test{}

	title, err := getTitle(lines)
	if err != nil {
		return test, err
	}
	test.Title = title

	strict, err := isStrict(lines[0])
	if err != nil {
		return test, err
	}
	test.Strict = strict

	whenBlock, err := getBlock(lines, BlockTypeWhen, false)
	if err != nil {
		return test, err
	}
	test.Script = strings.Join(whenBlock, "\n")

	thenBlock, err := getBlock(lines, BlockTypeThen, test.Strict)
	if err != nil {
		return test, err
	}
	test.ExpectedOutput = strings.Join(thenBlock, "\n")

	return test, nil
}

func getTitle(lines []string) (string, error) {
	if !lineIsTestBlockHeader(lines[0]) {
		return "", errors.New("Test Block Invalid")
	}
	title := lines[0][len(BlockTypeTest):]
	return strings.TrimSpace(title), nil
}

func isStrict(line string) (bool, error) {
	if !lineIsTestBlockHeader(line) {
		return false, errors.New("Test Block Invalid")
	}
	start := len(BlockTypeTest) + 1
	stop := len(BlockTypeTest) + 1 + len(StrictFlag)
	return line[start:stop] == StrictFlag, nil
}

// get the first set of \t prefixed lines within the specified block type
func getBlock(lines []string, blockType string, strict bool) ([]string, error) {
	buffer := make([]string, 0)
	inBlock := false
	capture := false
	lines = append(lines, BlockPrefix)
	for _, line := range lines {
		if !inBlock {
			if lineIsHeader(line, blockType) {
				inBlock = true
			}
		} else {
			if lineIsBlockBody(line) {
				capture = true
			}
			if capture {
				if lineIsBlockBody(line) || (!strict && line == "") {
					cleanLine := line
					if len(line) > 0 {
						cleanLine = line[1:]
					}
					buffer = append(buffer, cleanLine)
				} else {
					return buffer, nil
				}
			}

		}
	}
	return buffer, errors.New("Bad format: expected block missing")
}

// func lineIsBlockHeader(line string) bool {
// 	return lineIsHeader(line, BlockPrefix)
// }

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
