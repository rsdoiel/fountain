//
// fountain2html converts a Fountain File into an HTML fragement suitable for including
// like a scrippet.
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
	description = `fountain2html is a command line program that reads an fountain document and writes out HTML.
`

	examples = `Convert a *screenplay.fountain* to *screenplay.html*.

    fountain2html -i screenplay.foutnain -o screenplay.html

Or alternatively

    cat screenplay.fountain | foutnain2html > screenplay.html
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
	asHTMLPage       bool
	includeInlineCSS bool
	includeCSS       bool
	width            int
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
	app.BoolVar(&asHTMLPage, "html-page", false, "If true output an HTML page otherwise a fragement")
	app.BoolVar(&includeInlineCSS, "inline-css", false, "Add inline CSS")
	app.BoolVar(&includeCSS, "css", false, "Add link for CSS")
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
	// Override defaults
	fountain.AsHTMLPage = asHTMLPage
	fountain.MaxWidth = width
	fountain.InlineCSS = includeInlineCSS
	fountain.CSS = includeCSS
	// Parse  input and render screenplay
	screenplay, err := fountain.Run(src)
	cli.OnError(app.Eout, err, quiet)

	//and then render as a string
	if newLine {
		fmt.Fprintf(app.Out, "%s\n", screenplay)
	} else {
		fmt.Fprintf(app.Out, "%s", screenplay)
	}
}
