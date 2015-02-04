### test: this isn't right
#### when:
	echo "Foo"

#### then:
	Fool

### test: failed when clauses should cause an error
#### when:
	ls not_a_file

#### then:
	nothing here matters because the command won't work
