test: clean
	go test -v ./test

deps:
	go get -t

dist:
	go get github.com/mitchellh/gox
	gox -build-toolchain
	gox -output "dist/{{.Dir}}.{{.OS}}-{{.Arch}}"

clean:
	rm -r test/test_*

.PHONY: test deps clean

