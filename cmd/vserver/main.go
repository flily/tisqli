package main

import (
	"flag"
	"os"

	"github.com/flily/tisqli/vulnerable/server"
)

func main() {
	set := flag.NewFlagSet("vserver", flag.ExitOnError)
	args := os.Args[1:]

	server.Main(set, args)
}
