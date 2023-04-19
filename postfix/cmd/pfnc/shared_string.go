package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lsejx/go-notation/postfix/cmd"
)

const (
	helpOptionUnix = "-h"
	helpOptionGnu  = "--help"
	i64ModeId      = "i"
	u64ModeId      = "u"
)

func getCmdName() string {
	return filepath.Base(os.Args[0])
}

var helpSuggestionMessage = func() string {
	const body = "[" + helpOptionUnix + " | " + helpOptionGnu + "]" + " to print info"
	return fmt.Sprintf("%s %s", getCmdName(), body)
}()

var helpMessage = func() string {
	const body = `[type | option] [operand | operator] ...

type
    ` + i64ModeId + `    64bit signed int
    ` + u64ModeId + `    64bit unsigned int
    if no type selected, 64bit float

option
    ` + helpOptionUnix + `, ` + helpOptionGnu + `    print info

valid identifier
    64bit float mode           ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `  ` + cmd.Mod + `  ` + cmd.Pi + `  
    64bit signed int mode      ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `(rhs must be positive)  ` + cmd.Mod + `  ` + cmd.Max + `
    64bit unsigned int mode    ` + cmd.Add + `  ` + cmd.Sub + `  ` + cmd.Mul + `  ` + cmd.Div + `  ` + cmd.Pow + `  ` + cmd.Mod + `  ` + cmd.Max + `

SI
    k, M, G    available for all type
    m, u, n    available for only float mode`

	return fmt.Sprintf("%s %s", getCmdName(), body)
}()
