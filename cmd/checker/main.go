package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/flily/tisqli/tisqli/checker"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	detected, all := 0, 0

	for {
		fmt.Printf("> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		payload := strings.TrimSpace(string(line))
		if len(payload) <= 0 || strings.Index(payload, "#") == 0 {
			continue
		}

		all++
		result := checker.OnPartial(payload)
		fmt.Printf("[%v]: %s\n", result, payload)
		if result {
			detected++
		}
	}

	fmt.Printf("Detected: %d/%d  %8f%%\n", detected, all, 100*float64(detected)/float64(all))
}
