package cmd

import (
	"math"
	"strconv"
	"strings"
)

var f64Operations = map[string]func(float64, float64) (float64, error){
	Add: func(f1, f2 float64) (float64, error) {
		return f1 + f2, nil
	},
	Sub: func(f1, f2 float64) (float64, error) {
		return f1 - f2, nil
	},
	Mul: func(f1, f2 float64) (float64, error) {
		return f1 * f2, nil
	},
	Div: func(f1, f2 float64) (float64, error) {
		if f2 == 0.0 {
			return 0.0, ErrDivideByZero
		}
		return f1 / f2, nil
	},
	Pow: func(f1, f2 float64) (float64, error) {
		return math.Pow(f1, f2), nil
	},
	Mod: func(f1, f2 float64) (float64, error) {
		return math.Mod(f1, f2), nil
	},
}

func parseF64(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", "_")
	n, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return n, nil
	}

	switch s {
	case Pi:
		return math.Pi, nil
	}

	var si float64 = 1.0

	runes := []rune(s)
	l := len(runes)
	switch runes[l-1] {
	case Kilo:
		si = 1_000.0
	case Mega:
		si = 1_000_000.0
	case Giga:
		si = 1_000_000_000.0
	case Milli:
		si = 0.001
	case Micro:
		si = 0.00_000_1
	case Nano:
		si = 0.00_000_000_1
	default:
		return 0.0, newErrInvalidId(s)
	}

	coefficient, err := strconv.ParseFloat(string(runes[:l-1]), 64)
	if err != nil {
		return 0.0, newErrInvalidId(s)
	}

	return coefficient * si, nil
}

func formatF64(n float64) (string, error) {
	return strconv.FormatFloat(n, 'f', -1, 64), nil
}

var F64Runner = newRunner(f64Operations, parseF64, formatF64)
