package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lsejx/go-notation/postfix/cmd"
)

const (
	helpOptionUnix    = "-h"
	helpOptionGnu     = "--help"
	i64ModeOptionUnix = "-i"
	i64ModeOptionGnu  = "--int"
	u64ModeOptionUnix = "-u"
	u64ModeOptionGnu  = "--uint"
)

func getCmdName() string {
	return filepath.Base(os.Args[0])
}

var helpSuggestionMessage = func() string {
	const body = "[" + helpOptionUnix + " | " + helpOptionGnu + "]" + " to print info"
	return fmt.Sprintf("%s %s", getCmdName(), body)
}()

var helpMessage = func() string {
	const body = `[option] [operand | operator] ...

the default is 64bit float mode

option
    ` + helpOptionUnix + `, ` + helpOptionGnu + `    print info
    ` + i64ModeOptionUnix + `, ` + i64ModeOptionGnu + `     64bit signed int mode
    ` + u64ModeOptionUnix + `, ` + u64ModeOptionGnu + `    64bit unsigned int mode

valid identifier
    64bit float mode           ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `  ` + cmd.Mod + `  ` + cmd.Pi + `  
    64bit signed int mode      ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `(rhs must be positive)  ` + cmd.Mod + `  ` + cmd.Max + `  ` + cmd.Min + `
    64bit unsigned int mode    ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `  ` + cmd.Mod + `  ` + cmd.Max + `

SI
    k, M, G    available for all type
    m, u, n    available for only float mode`

	return fmt.Sprintf("%s %s", getCmdName(), body)
}()
