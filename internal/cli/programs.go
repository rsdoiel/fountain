package cli

import (
	"fmt"
	"io"

	"github.com/rsdoiel/fountain"
)

/** Fountain2html is fountain2html: it reads a Fountain document and writes
 * HTML, a fragment or a page.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Fountain2html("fountain2html", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Fountain2html(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	var page, inlineCSS, linkCSS bool
	var css string
	fset := newFlagSet(appName, &o)
	fset.BoolVar(&page, "page", false, "output an HTML page, not a fragment")
	fset.BoolVar(&inlineCSS, "inline-css", false, "add inline CSS")
	fset.BoolVar(&linkCSS, "link-css", false, "add a link to CSS")
	fset.StringVar(&css, "css", "fountain.css", "include a custom CSS file")
	if code, done := start(appName, Fountain2htmlHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	doc, code := parse(appName, &o, stdin, stderr)
	if code != ExitOK {
		return code
	}
	fountain.AsHTMLPage = page
	fountain.MaxWidth = o.width
	fountain.InlineCSS = inlineCSS
	fountain.LinkCSS = linkCSS
	fountain.CSS = css
	return write(appName, &o, []byte(doc.ToHTML()), stdout, stderr)
}

/** Fountain2json is fountain2json: it reads a Fountain document and writes it
 * as JSON.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Fountain2json("fountain2json", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Fountain2json(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	var pretty bool
	fset := newFlagSet(appName, &o)
	fset.BoolVar(&pretty, "pretty", false, "pretty print the JSON output")
	if code, done := start(appName, Fountain2jsonHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	doc, code := parse(appName, &o, stdin, stderr)
	if code != ExitOK {
		return code
	}
	fountain.MaxWidth = o.width
	fountain.PrettyPrint = pretty
	src, err := doc.ToJSON()
	if err != nil {
		return fail(appName, &o, stderr, ExitInternal, "%v", err)
	}
	return write(appName, &o, src, stdout, stderr)
}

/** Fountainfmt is fountainfmt: it reads a Fountain document and pretty prints
 * it, or with -debug lists its elements.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Fountainfmt("fountainfmt", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Fountainfmt(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	var debug, section, synopsis, notes bool
	fset := newFlagSet(appName, &o)
	fset.BoolVar(&debug, "debug", false, "display type and element content")
	fset.BoolVar(&section, "section", false, "include sections in output")
	fset.BoolVar(&synopsis, "synopsis", false, "include synopsis in output")
	fset.BoolVar(&notes, "notes", false, "include notes in output")
	if code, done := start(appName, FountainfmtHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	fountain.MaxWidth = o.width
	fountain.ShowSection = section
	fountain.ShowSynopsis = synopsis
	fountain.ShowNotes = notes
	doc, code := parse(appName, &o, stdin, stderr)
	if code != ExitOK {
		return code
	}
	if debug {
		var list []byte
		for i, element := range doc.Elements {
			list = fmt.Appendf(list, "%4d %02d %q\n", i, element.Type, element.Content)
		}
		o.newline = false
		return write(appName, &o, list, stdout, stderr)
	}
	return write(appName, &o, []byte(doc.String()), stdout, stderr)
}
