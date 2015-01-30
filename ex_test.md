I can say antyhing here

### test: echo should work
- take notes here
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
	

### test: "strict" flag enables testing trailing white space
#### when:
	echo -e "Foo\n\n\tBar"


#### then:
	Foo
	
		Bar
	# strict

## and start a new section over here
ha!
