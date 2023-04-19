package cmd

import (
	"math"
	"strconv"
	"strings"
)

var i64Operations = map[string]func(int64, int64) (int64, error){
	Add: func(i1, i2 int64) (int64, error) {
		return i1 + i2, nil
	},
	Sub: func(i1, i2 int64) (int64, error) {
		return i1 - i2, nil
	},
	Mul: func(i1, i2 int64) (int64, error) {
		return i1 * i2, nil
	},
	Div: func(i1, i2 int64) (int64, error) {
		if i2 == 0 {
			return 0, ErrDivideByZero
		}
		return i1 / i2, nil
	},
	Pow: func(i1, i2 int64) (int64, error) {
		if i2 < 0 {
			return 0, ErrPowNegativeInt
		}
		var result int64 = 1
		for i := int64(0); i < i2; i++ {
			result *= i1
		}
		return result, nil
	},
	Mod: func(i1, i2 int64) (int64, error) {
		if i2 == 0 {
			return 0, ErrDivideByZero
		}
		return i1 % i2, nil
	},
}

func parseI64(s string) (int64, error) {
	s = strings.ReplaceAll(s, ",", "_")
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return n, nil
	}

	switch s {
	case Max:
		return math.MaxInt64, nil
	}

	var si int64 = 1

	runeSlice := []rune(s)
	l := len(runeSlice)
	switch runeSlice[l-1] {
	case Kilo:
		si = 1_000
	case Mega:
		si = 1_000_000
	case Giga:
		si = 1_000_000_000
	default:
		return 0, newErrInvalidId(s)
	}

	coefficient, err := strconv.ParseInt(string(runeSlice[:l-1]), 10, 64)
	if err != nil {
		return 0, newErrInvalidId(s)
	}

	return coefficient * si, nil
}

func formatI64(n int64) (string, error) {
	return strconv.FormatInt(n, 10), nil
}

var I64Runner = newRunner(i64Operations, parseI64, formatI64)
