#
# Simple Makefile
#
PROJECT = fountain

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

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

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

