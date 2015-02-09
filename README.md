[![Build Status](https://drone.io/github.com/benschw/cli-unit/status.png)](https://drone.io/github.com/benschw/cli-unit/latest)

## cli-unit

unit test your command line interfaces and use your tests as usage documentation

* [Download](http://txt.fliglio.com/cli-unit/)


## Install

- add `cli-unitw.sh` to your project
- add `.fli-unit` to your `.gitignore` file
- run your test files: `./cli-unitw.sh *_test.md`
 

## Test Files

- `### test:` signals the start of a new test. Also sets the test description
- `#### when:` starts the block where you can define your script usage
- `#### then:` starts the block where you define the expected output


## Examples

### test: Vanilla test run
#### when:

	./cli-unit ./ex_test.md 

#### then:

	Pass (4/4 tests successful)

### test: Optionally get verbose output
#### when:

	./cli-unit -v ./ex_test.md 

#### then:

	--- OK: echo should work
	--- OK: echo -e should preserve special chars
	--- OK: (strict) strict flag enables testing trailing white space
	--- OK: pipes should work too
	Pass (4/4 tests successful)

### test: failed tests should show what the problem is (also handle expected error)
#### when:

	./cli-unit failures_ex.md || true


#### then:

	--- FAIL: this isn't right
	"Fool" != "Foo"
	--- FAIL: failed when clauses should cause an error
	exit status 2: ls: cannot access not_a_file: No such file or directory
	Fail (0/2 tests successful)

## Formatting Examples

you can have multiple blocks of tests in your markdown files (we didn't have to jump directly from "Examples" to "Formatting Examples")

### test: check some output
You can put whatever you want here
#### when:
or here

	echo -e "foo"

- and not
- just text

#### then:

or here

	foo

or here...

### test: check some more output
when not using `(strict)`, empty lines count as part of your `when` and `then` blocks
and trailing whitespace characters aren't compared

#### when:

	echo -e "foo\n\n\nbar\n\n"

#### then:

	foo


	bar


### test: (strict) check some strict output
when using `(strict)`, empty lines terminate your `when` and `then` blocksand trailing
tabbed lines are compared

#### when:

	echo -e "foo\nbar\n"

#### then:
strict mode is hard to set up right if you're using some editors like sublime text which will do a bunch of auto trimming.


	foo
	bar
	
	



	this is not part of the test
