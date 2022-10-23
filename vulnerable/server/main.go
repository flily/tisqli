package server

import (
	"flag"
)

func Main(set *flag.FlagSet, args []string) {
	port := set.Int("port", 8000, "port to listen on")
	dsn := set.String("dsn", "root@/tisqli", "data source name")
	partial := set.Bool("partial", false, "enable partial SQL injection checking")
	full := set.Bool("full", false, "enable full SQL injection checking")
	_ = set.Parse(args)

	server, err := NewServer(*dsn, *partial, *full)
	if err != nil {
		panic(err)
	}

	err = server.Init()
	if err != nil {
		panic(err)
	}

	err = server.Serve(*port)
	if err != nil {
		panic(err)
	}
}
