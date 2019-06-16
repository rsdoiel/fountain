#
# Simple Makefile
#
PROJECT = fountain

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build: fountain.go cmd/fountainfmt/fountainfmt.go
	go build -o bin/fountainfmt$(EXT) cmd/fountainfmt/fountainfmt.go

test:
	go test

man: build
	mkdir -p man/man1
	bin/fountainfmt -generate-manpage | nroff -Tutf8 -man > man/man1/fountainfmt.1

install:
	env GOBIN=$(HOME)/bin go install cmd/fountainfmt/fountainfmt.go

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fountainfmt.exe cmd/fountainfmt/fountainfmt.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/linux-arm64

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	./mk-website.py

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

publish:
	./mk-website.py
	./publish.bash

