package main

import (
	"reflect"
	"strings"
	"testing"
)

func NewStubFileParser(content string) *TestFileParser {
	lines := strings.Split(content, "\n")

	return &TestFileParser{
		lines: lines,
		idx:   0,
	}
}

func Test_fileParser_with_vanilla_content(t *testing.T) {
	// given
	content := `
### test: echo should work
#### when:
	echo "Foo"
#### then:
	Foo
`

	expected := &Test{
		Title:          "echo should work",
		Script:         `echo "Foo"`,
		ExpectedOutput: "Foo\n",
	}

	parser := NewStubFileParser(content)

	// when
	test1, err1 := parser.NextTest()
	test2, err2 := parser.NextTest()

	// then
	if err1 != nil {
		t.Errorf("Unexpected Error: %s", err1)
	}
	if err2 != nil {
		t.Errorf("Unexpected Error: %s", err2)
	}
	if test2 != nil {
		t.Errorf("second call of NextTest() should return null; found %+v", test2)
	}

	if !reflect.DeepEqual(test1, expected) {
		t.Errorf("Expected/Found:\n%+v\n%+v", expected, test1)
	}
}

func Test_fileParser_with_surrounding_markdown(t *testing.T) {
	// given
	content := `
# Hello!
- this
- shouldn't
- matter
### test: echo should work
- this
- is
	cool
	too
#### when:
hello

	echo "Foo"
hey!
#### then:

	Foo

asd
	junk
sup

### test: echo should work
#### when:
	echo "Foo"
#### then:

	Foo

sup

## new header
foo
`

	expected := &Test{
		Title:          "echo should work",
		Script:         `echo "Foo"`,
		ExpectedOutput: "Foo\n",
	}

	parser := NewStubFileParser(content)

	// when
	test1, _ := parser.NextTest()
	test2, _ := parser.NextTest()
	test3, _ := parser.NextTest()

	// then
	if test3 != nil {
		t.Errorf("second call of NextTest() should return null; found %+v", test2)
	}

	if !reflect.DeepEqual(test1, expected) {
		t.Errorf("Expected/Found:\n%+v\n%+v", expected, test1)
	}
	if !reflect.DeepEqual(test2, expected) {
		t.Errorf("Expected/Found:\n%+v\n%+v", expected, test2)
	}
}

func Test_fileParser_default_comparison_should_trim_subjects(t *testing.T) {
	// given

	// watch out: sublime likes to trim whitespace, the line after Foo in
	// the "then" block *might* start with a tab. this only matters with "strict"
	// usage though
	content := `
### test: echo should work
#### when:
	echo "Foo"
#### then:

	Foo
	

`

	expectedOutput := "Foo"

	parser := NewStubFileParser(content)

	// when
	test, _ := parser.NextTest()

	// then

	if test.ExpectedOutput == expectedOutput {
		t.Errorf("shouldn't match: Expected/Found:\n%+v\n%+v", test.ExpectedOutput, expectedOutput)
	}

	if test.GetExpectedOutput() != expectedOutput {
		t.Errorf("Expected/Found:\n'%s':'%s'", test.GetExpectedOutput(), expectedOutput)
	}
}

func Test_fileParser_strict_comparison_should_not_trim_subjects(t *testing.T) {
	// given

	// watch out: sublime likes to trim whitespace, the line after Foo in
	// the "then" block *might* start with a tab. this only matters with "strict"
	// usage though
	content := `
### test: (strict) echo should work
#### when:
	echo "Foo"
#### then:

	Foo
	
	

`

	expectedOutput := "Foo\n\n"

	parser := NewStubFileParser(content)

	// when
	test, _ := parser.NextTest()

	// then

	if test.ExpectedOutput != expectedOutput {
		t.Errorf("shouldn't match: Expected/Found:\n%+v\n%+v", test.ExpectedOutput, expectedOutput)
	}

	if test.GetExpectedOutput() != expectedOutput {
		t.Errorf("Expected/Found:\n'%s':'%s'", test.GetExpectedOutput(), expectedOutput)
	}
}
