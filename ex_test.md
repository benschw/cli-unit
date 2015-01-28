### test: echo should work
#### when:
	echo "Foo"


#### then:
	Foo

### test: echo -e should preserve special chars
#### when:
	echo -e "Foo\n\tBar"


#### then:
	Foo
		Bar
