// fountainfmt is part of fountain, a Go package for Fountain screenplays. Run fountainfmt -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/fountain/internal/cli"
)

func main() {
	os.Exit(cli.Fountainfmt(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
