.PHONY=test

test:
	go test

bench:
	go test -bench .
