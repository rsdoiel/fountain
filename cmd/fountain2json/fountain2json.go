//
// fountain2json parses a fountain file and returns a JSON representation.
//
// fountain is a package encoding/decoding fountain formatted screenplays.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2019, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	// My packages
	"github.com/rsdoiel/fountain"
)

var (

	helpText = `%{app_name}(1) | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a command line program that reads an fountain document and returns a JSON representation of it.

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i
: read from filename

-o
: write to filename

-newline
: add a trailing newline

-width
: set the width of the text

-pretty
: pretty print the output


# EXAMPLES

Render *screenplay.fountain* as *screenplay.json*.

~~~
{app_name} -i screenplay.fountain -o screenplay.json
~~~

Or alternatively

~~~
    cat screenplay.fountain | {app_name} > screenplay.json
~~~

`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	newLine          bool
	quiet            bool
	inputFName       string
	outputFName      string

	// App Option
	width       int
	prettyPrint bool
)

func main() {
	appName := path.Base(os.Args[0])
	// NOTE: the following are set with version.go is generted
	version := fountain.Version
	releaseDate := fountain.ReleaseDate
	releaseHash := fountain.ReleaseHash
	fmtHelp := fountain.FmtHelp

	// Standard Options
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&newLine, "newline", true, "add a trailing newline")
	flag.BoolVar(&quiet, "quiet", false, "suppress error messages")
	flag.StringVar(&inputFName, "i", "", "set the input filename")
	flag.StringVar(&outputFName, "o", "", "set the output filename")

	// App Option
	flag.BoolVar(&prettyPrint, "pretty", false, "pretty print the JSON output")
	flag.IntVar(&width, "width", 65, "set the width for the text")

	// Parse environment and options
	flag.Parse()

	// Setup IO
	var err error

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if inputFName != "" {
		in, err = os.Open(inputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer in.Close()
	}
	
	if outputFName != "" {
		out, err = os.Create(outputFName)
		if err != nil {
			fmt.Fprintf(eout, "%s\n", err)
			os.Exit(1)
		}
		defer out.Close()
	}

	// Process options
	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", fountain.LicenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s %s\n", appName, version, releaseHash)
		os.Exit(0)
	}

	// ReadAll of input
	src, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	// Parse input
	screenplay, err := fountain.Parse(src)
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	fountain.PrettyPrint = prettyPrint
	src, err = screenplay.ToJSON()
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(out, "%s", src)
	if newLine {
		fmt.Fprintln(out)
	}
}
