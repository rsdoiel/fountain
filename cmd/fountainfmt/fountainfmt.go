//
// fountainfmt pretty prints a fountain file.
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"

	// My packages
	"github.com/rsdoiel/fountain"
)

var (
	description = `fountainfmt is a command line program that reads an fountain document and pretty prints it.
`

	examples = `Pretty print *screenplay.txt* saving it as *screenplay.fountain*.

    fountainfmt -i screenplay.txt -o screenplay.fountain

Or alternatively

    cat screenplay.txt | foutnainfmt > screenplay.fountain
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	generateMarkdown bool
	generateManPage  bool
	newLine          bool
	quiet            bool
	inputFName       string
	outputFName      string

	// App Option
	width int
	debug bool
)

func main() {
	app := cli.NewCli(fountain.Version)

	// Add Help
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&newLine, "nl,newline", true, "add a trailing newline")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.StringVar(&inputFName, "i,input", "", "set the input filename")
	app.StringVar(&outputFName, "o,output", "", "set the output filename")

	// App Option
	app.BoolVar(&debug, "debug", false, "display type and element content")
	app.IntVar(&width, "w,width", 65, "set the width for the text")

	// Parse environment and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error
	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	// ReadAll of input
	src, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)
	// Parse input
	screenplay, err := fountain.Parse(src)
	cli.OnError(app.Eout, err, quiet)

	//and then render as a string
	fountain.MaxWidth = width
	if debug {
		for i, element := range screenplay.Elements {
			fmt.Fprintf(app.Out, "%4d %02d %q\n", i, element.Type, element.Content)
		}
		os.Exit(0)
	}
	if newLine {
		fmt.Fprintf(app.Out, "%s\n", screenplay.String())
	} else {
		fmt.Fprintf(app.Out, "%s", screenplay.String())
	}
}
