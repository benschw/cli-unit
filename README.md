
## cli-unit

unit test your command line interfaces


## Install

- add `cli-unitw.sh` to your project
- add `.fli-unit` to your `.gitignore` file
- run your test files: `./cli-unitw.sh *_test.md`
 

## Test Files

- `### test:` signals the start of a new test. Also sets the test description
- `### shell:` starts the block where you can define your script usage
- `### output:` starts the block where you define the expected output

	### test: -json flag should make output be in json
	### shell:
	echo "Foo"


	### output:
	Foo

	### test: should target array element
	#### shell:
	echo "Bar"


	#output
	Bar
