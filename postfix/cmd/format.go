package cmd

import (
	"errors"
	"fmt"
	"strings"
)

func insertComma(s string) string {
	out := make([]byte, len(s)+(len(s)-1)/3)
	slast := len(s) - 1
	olast := len(out) - 1

	cursor := 0
	for i := 0; i < len(s); i++ {
		if i%3 == 0 && i != 0 {
			out[olast-cursor] = ','
			cursor++
		}
		out[olast-cursor] = s[slast-i]
		cursor++
	}
	return string(out)
}

func format(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	var signEnd uint = 0
	switch []rune(s)[0] {
	case '-', '+':
		if len(s) == 1 {
			return "", errors.New("only a sign, there is not number")
		}
		signEnd = 1
	}

	for _, b := range []byte(s) {
		if b != '+' && b != '-' && b != '.' && !(b >= '0' && b <= '9') {
			return "", fmt.Errorf("%c is invalid", b)
		}
	}

	periodSplitted := strings.Split(s, ".")
	formatted := periodSplitted[0][:signEnd] + insertComma(periodSplitted[0][signEnd:])
	if len(periodSplitted) == 1 {
		return formatted, nil
	}
	if len(periodSplitted) == 2 {
		return formatted + "." + periodSplitted[1], nil
	}

	return "", errors.New("there are two periods")
}
