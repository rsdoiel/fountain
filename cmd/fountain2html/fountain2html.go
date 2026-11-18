// fountain2html is part of fountain, a Go package for Fountain screenplays. Run fountain2html -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/fountain/internal/cli"
)

func main() {
	os.Exit(cli.Fountain2html(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
