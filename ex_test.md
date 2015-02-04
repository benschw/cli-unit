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
	

### test: (strict) strict flag enables testing trailing white space
#### when:
	echo -e "Foo\n\n\tBar"


#### then: 
	Foo
	
		Bar
	



### test: pipes should work too
#### when:
	echo -e "hello\nworld" | grep -v hello


#### then:
	world
	


## and start a new section over here
ha!
