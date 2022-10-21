package checker

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func doPartial(payload string, isDebug bool) bool {
	checker := DefaultPartialChecker()
	result := checker.Check(payload)
	fmt.Printf("[ %5v ] %s\n", result.IsInjection(), payload)
	for _, t := range result.Results {
		fmt.Printf("  - [%s] %s\n", t.Reason, t.SQLInColour())

		if isDebug {
			fmt.Printf("%s\n", t.Err)

			for _, a := range t.AstCorrect {
				a.PrintTree(0, "", true)
			}

			for _, a := range t.AstPartial {
				a.PrintTree(0, "", true)
			}
		}
	}

	return result.IsInjection()
}

func Main(set *flag.FlagSet, args []string) {
	isFull := set.Bool("full", false, "check for full SQL statements")
	isDebug := set.Bool("debug", false, "debug mode")
	_ = set.Parse(args)

	reader := bufio.NewReader(os.Stdin)
	detected, all := 0, 0
	stat, _ := os.Stdin.Stat()
	istty := (stat.Mode() & os.ModeCharDevice) != 0

	for {
		if istty {
			fmt.Printf("> ")
		}

		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		payload := string(line)
		if len(payload) <= 0 || strings.Index(payload, "#") == 0 {
			continue
		}

		all++
		result := false
		if !*isFull {
			result = doPartial(payload, *isDebug)
		}

		if result {
			detected++
		}
	}

	fmt.Printf("Detected: %d/%d  %8f%%\n", detected, all, 100*float64(detected)/float64(all))
}
