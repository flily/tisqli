package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/flily/tisqli/syntax"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		parser := syntax.NewParser()
		nodes, warns, err := parser.Parse(string(line))
		if err != nil {
			fmt.Printf("parse error\n")
			fmt.Printf("%s\n", err)
			continue
		}

		for _, warn := range warns {
			fmt.Printf("warning: %s\n", warn)
		}

		fmt.Printf("%d SQLs found\n", len(nodes))
		fmt.Printf("%s\n", strings.Repeat("-", 80))
		for _, node := range nodes {
			node.VerifyConstant()
			node.PrintTree(0, "", true)
			fmt.Printf("%s\n", strings.Repeat("-", 80))
		}
	}
}
