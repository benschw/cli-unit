package main

import (
	"testing"
)

func Test_Test_Run_should_fail_if_error(t *testing.T) {
	// given
	content := `
### test: ls not existing file should fail
#### when:
	ls not_a_file
#### then:
	Foo
`
	expectedError := "exit status 2: ls: cannot access not_a_file: No such file or directory"
	test := &Test{
		Script: `ls not_a_file`,
	}

	parser := NewStubFileParser(content)

	// when
	test, err := parser.NextTest()
	test.Run()

	// then
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}

	if test.Pass() {
		t.Error("Test should fail if script fails")
	}
	if test.GetFailMessage() != expectedError {
		t.Errorf("Expected/Found:\n'%s'\n'%s'", expectedError, test.GetFailMessage())
	}
}
