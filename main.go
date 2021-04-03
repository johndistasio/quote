package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	noFlag     = flag.Bool("no", false, "don't quote")
	nullFlag   = flag.Bool("0", false, "output null instead of newline")
	singleFlag = flag.Bool("single", false, "use single quotes")
	spaceFlag  = flag.Bool("space", false, "escape spaces")
)

func main() {
	flag.Parse()

	quote := `"`

	if *singleFlag {
		quote = `'`
	}

	terminator := "\n"

	if *nullFlag {
		terminator = "\u0000"
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if !*noFlag {
			line = strings.ReplaceAll(line, quote, `\`+quote)
		}

		if *spaceFlag {
			line = strings.ReplaceAll(line, ` `, `\ `)
		}

		fmt.Printf("%s%s%s%s", quote, line, quote, terminator)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "quote: %s\n", err)
		os.Exit(1)
	}
}
