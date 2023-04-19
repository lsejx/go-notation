package cmd

import (
	"math"
	"strconv"
	"strings"
)

var u64Operations = map[string]func(uint64, uint64) (uint64, error){
	Add: func(u1, u2 uint64) (uint64, error) {
		return u1 + u2, nil
	},
	Sub: func(u1, u2 uint64) (uint64, error) {
		return u1 - u2, nil
	},
	Mul: func(u1, u2 uint64) (uint64, error) {
		return u1 * u2, nil
	},
	Div: func(u1, u2 uint64) (uint64, error) {
		if u2 == 0 {
			return 0, ErrDivideByZero
		}
		return u1 / u2, nil
	},
	Pow: func(u1, u2 uint64) (uint64, error) {
		var result uint64 = 1
		for i := uint64(0); i < u2; i++ {
			result *= u1
		}
		return result, nil
	},
	Mod: func(u1, u2 uint64) (uint64, error) {
		if u2 == 0 {
			return 0, ErrDivideByZero
		}
		return u1 % u2, nil
	},
}

func parseU64(s string) (uint64, error) {
	s = strings.ReplaceAll(s, ",", "_")
	n, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return n, nil
	}

	switch s {
	case Max:
		return math.MaxUint64, nil
	}

	var si uint64 = 1

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

	coefficient, err := strconv.ParseUint(string(runeSlice[:l-1]), 10, 64)
	if err != nil {
		return 0, newErrInvalidId(s)
	}

	return coefficient * si, nil
}

func formatU64(n uint64) (string, error) {
	return strconv.FormatUint(n, 10), nil
}

var U64Runner = newRunner(u64Operations, parseU64, formatU64)
