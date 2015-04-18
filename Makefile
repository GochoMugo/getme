test:
	go test -v ./test

deps:
	go get -t

dist:
	go get github.com/mitchellh/gox
	gox -build-toolchain
	gox -output "dist/{{.Dir}}.{{.OS}}-{{.Arch}}"

.PHONY: test

