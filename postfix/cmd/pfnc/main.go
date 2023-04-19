package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lsejx/go-notation/postfix"
	"github.com/lsejx/go-notation/postfix/cmd"
)

func eprintf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func eprintln(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func main() {
	if len(os.Args) == 1 {
		eprintln("no formula")
		eprintln(helpSuggestionMessage)
		os.Exit(1)
	}
	args := os.Args[1:]

	var (
		result string
		err    error
	)

	switch args[0] {
	case helpOptionUnix, helpOptionGnu:
		fmt.Println(helpMessage)
		os.Exit(0)
	case i64ModeId:
		result, err = cmd.I64Runner.Run(args[1:])
	case u64ModeId:
		result, err = cmd.U64Runner.Run(args[1:])
	default:
		result, err = cmd.F64Runner.Run(args)
	}

	if err != nil {
		if errors.Is(err, postfix.ErrNoFormula) {
			eprintf("development error: %v\n", err)
			os.Exit(1)
		}
		eprintf("error: %v\n", err)
		eprintln(helpSuggestionMessage)
		os.Exit(1)
	}

	fmt.Println(result)
}
