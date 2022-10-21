package main

import (
	"flag"
	"os"

	"github.com/flily/tisqli/tisqli/checker"
)

func main() {
	args := os.Args[1:]
	set := flag.NewFlagSet("checker", flag.ExitOnError)
	checker.Main(set, args)
}
