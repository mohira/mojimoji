package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	ExitCodeOK         = 0
	ExitCodeSomeError  = 1
	ExitCodeParseError = 2
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}
func (c CLI) Run(args []string) int {
	flags := flag.NewFlagSet("mojimoji", flag.ContinueOnError)

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseError
	}

	filename := flags.Arg(0)

	f, err := os.Open(filename)
	if err != nil {
		if _, err := fmt.Fprint(c.errStream, err); err != nil {
			return ExitCodeSomeError
		}
		return ExitCodeSomeError
	}

	scanner := bufio.NewScanner(f)

	no := 0
	for scanner.Scan() {
		no++
		line := scanner.Text()

		ignore := line == "" || strings.HasPrefix(line, "//")

		if ignore {
			continue
		}

		for _, statement := range limitOverStatements(line) {
			msg := fmt.Sprintf("L%d(%d): %s\n", no, utf8.RuneCountInString(statement), statement)

			if _, err := fmt.Fprint(c.outStream, msg); err != nil {
				return ExitCodeSomeError
			}
		}
	}

	return ExitCodeOK
}

func limitOverStatements(s string) []string {
	const CharacterLimit int = 35

	var result []string

	for _, statement := range strings.Split(s, "。") {
		statement += "。"
		count := utf8.RuneCountInString(statement)
		if count > CharacterLimit {
			result = append(result, statement)
		}
	}

	return result
}
