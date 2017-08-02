package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ExitOnError)
	version := flags.Bool("version", false, "Print version and exit")
	flags.Usage = func() {
		fmt.Fprintf(cli.errStream, usage, Name)
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if *version {
		fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}

const usage = `Usage: %s [options]

Options:

  -version	Print the version of this application
`
