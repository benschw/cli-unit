### test: echo should work
#### when:
	echo "Foo"


#### then:
	Fool

### test: echo -e should preserve special chars
#### when:
	echo -e "Foo\n\tBar"


#### then:
	Food
		Bar
