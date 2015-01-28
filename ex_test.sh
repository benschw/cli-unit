### test: echo should work
### shell:
echo "Foo"


### output:
Foo

### test: echo -e should preserve special chars
### shell:
echo -e "Foo\n\tBar"


### output:
Foo
	Bar
