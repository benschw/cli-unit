
## cli-unit

unit test your command line interfaces


## Install

- add `cli-unitw.sh` to your project
- add `.fli-unit` to your `.gitignore` file
- run your test files: `./cli-unitw.sh *_test.md`
 

## Test Files

- `### test:` signals the start of a new test. Also sets the test description
- `#### when:` starts the block where you can define your script usage
- `#### then:` starts the block where you define the expected output


## suite: Example

### test: sucessful tests should "Pass"
#### when:

	./cli-unit ./ex_test.md 

#### then:

	Pass (2/2 tests successful)

### test: -v should add result of each test to output
#### when:

	./cli-unit -v ./ex_test.md 

#### then:

	--- OK: echo should work
	--- OK: echo -e should preserve special chars
	Pass (2/2 tests successful)

### test: failed tests should show what the problem is
#### when:

	./cli-unit failures_ex.md 

#### then:

	--- FAIL: echo should work
	"Fool" != "Foo\n"
	--- FAIL: echo -e should preserve special chars
	"Food\n\tBar" != "Foo\n\tBar\n"
	Fail (0/2 tests successful)
