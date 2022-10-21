package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/flily/tisqli/tisqli/checker"
)

type cliConfig struct {
	IsDebug         bool
	FilterPositives bool
	FilterNegatives bool
}

func doPartial(payload string, conf *cliConfig) (*checker.PartialResult, time.Duration) {
	timeStart := time.Now()
	checker := checker.DefaultPartialChecker()
	result := checker.Check(payload)
	timeFinish := time.Now()
	timeDuration := timeFinish.Sub(timeStart)

	if conf.FilterPositives && result.IsInjection() {
		return result, timeDuration
	}

	if conf.FilterNegatives && !result.IsInjection() {
		return result, timeDuration
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

	return result, timeDuration
}

func main() {
	conf := &cliConfig{}
	isFull := flag.Bool("full", false, "check for full SQL statements")
	flag.BoolVar(&conf.IsDebug, "debug", false, "debug mode")
	flag.BoolVar(&conf.FilterPositives, "negative", false, "filter positives")
	flag.BoolVar(&conf.FilterNegatives, "positive", false, "filter negatives")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	detected, all := 0, 0
	totalTime := time.Duration(0)

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
			totalTime += duration
		}

		if result {
			detected++
		}
	}

	fmt.Printf("Detected: %d/%d  %8f%%  avg=%s\n",
		detected, all,
		100*float64(detected)/float64(all),
		totalTime/time.Duration(all),
	)

	for t, c := range tc {
		fmt.Printf("[ %6d ] %s\n", c, t)
	}
}
