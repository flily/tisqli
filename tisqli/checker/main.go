package checker

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type cliConfig struct {
	IsDebug         bool
	FilterPositives bool
	FilterNegatives bool
}

func doPartial(payload string, conf *cliConfig) (*PartialResult, time.Duration) {
	time_start := time.Now()
	checker := DefaultPartialChecker()
	result := checker.Check(payload)
	time_end := time.Now()
	time_elapsed := time_end.Sub(time_start)

	if conf.FilterPositives && result.IsInjection() {
		return result, time_elapsed
	}

	if conf.FilterNegatives && !result.IsInjection() {
		return result, time_elapsed
	}

	fmt.Printf("[ %5v ] %s\n", result.IsInjection(), payload)
	for _, t := range result.Results {
		fmt.Printf("  - [%s] %s\n", t.Reason, t.SQLInColour())

		if conf.IsDebug {
			fmt.Printf("%s\n", t.Err)

			for _, a := range t.AstCorrect {
				a.PrintTree(0, "", true)
			}

			for _, a := range t.AstPartial {
				a.PrintTree(0, "", true)
			}
		}
	}

	return result, time_elapsed
}

func Main(set *flag.FlagSet, args []string) {
	conf := &cliConfig{}
	isFull := set.Bool("full", false, "check for full SQL statements")
	set.BoolVar(&conf.IsDebug, "debug", false, "debug mode")
	set.BoolVar(&conf.FilterPositives, "negative", false, "filter positives")
	set.BoolVar(&conf.FilterNegatives, "positive", false, "filter negatives")
	_ = set.Parse(args)

	reader := bufio.NewReader(os.Stdin)
	detected, all := 0, 0
	total_times := time.Duration(0)

	stat, _ := os.Stdin.Stat()
	istty := (stat.Mode() & os.ModeCharDevice) != 0
	tc := map[string]int{}

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
			r, duration := doPartial(payload, conf)
			result = r.IsInjection()
			for _, t := range r.Results {
				if t.IsInjection {
					tc[t.Template]++
				}
			}
			total_times += duration
		}

		if result {
			detected++
		}
	}

	fmt.Printf("Detected: %d/%d  %8f%%  avg=%s\n",
		detected, all,
		100*float64(detected)/float64(all),
		total_times/time.Duration(all),
	)

	for t, c := range tc {
		fmt.Printf("[ %6d ] %s\n", c, t)
	}
}
