package cmd

import (
	"errors"
	"fmt"

	"github.com/lsejx/go-notation/postfix"
)

// operator
const (
	Add = "+"
	Sub = "-"
	Mul = "x"
	Div = "/"
	Pow = "xx"
	Mod = "%"
)

// special constant Id
const (
	Pi  = "pi"
	Max = "max"
	Min = "min"
)

// SI prefix
const (
	Kilo  = 'k'
	Mega  = 'M'
	Giga  = 'G'
	Milli = 'm'
	Micro = 'u'
	Nano  = 'n'
)

var (
	ErrDivideByZero   = errors.New("divide/mod by zero")
	ErrPowNegativeInt = errors.New("cannot pow on negative integer")
)

var newErrInvalidId = func(ids ...string) error {
	if len(ids) == 0 {
		return nil
	}
	text := "invalid:"
	for _, id := range ids {
		text += " "
		text += id
	}
	return errors.New(text)
}

type Runner[T any] struct {
	calculator *postfix.Calculator[T]
	operations map[string]func(T, T) (T, error)
	parse      func(string) (T, error)
	format     func(T) (string, error)
}

func newRunner[T any](operations map[string]func(T, T) (T, error), parse func(string) (T, error), format func(T) (string, error)) Runner[T] {
	return Runner[T]{
		calculator: postfix.NewCalculator[T](2),
		operations: operations,
		parse:      parse,
		format:     format,
	}
}

func (runner Runner[T]) Run(formula []string) (string, error) {
	for _, elm := range formula {
		operation, isPresent := runner.operations[elm]
		if isPresent {
			_, err := runner.calculator.Operate(operation)
			if err != nil {
				return "", err
			}
			continue
		}

		n, err := runner.parse(elm)
		if err != nil {
			return "", err
		}
		runner.calculator.AppendOperand(n)
	}

	result, err := runner.calculator.Result()
	if err != nil {
		return "", err
	}

	s, err := runner.format(result)
	if err != nil {
		return "", err
	}

	s, err = format(s)
	if err != nil {
		return "", fmt.Errorf("development error: %v", err)
	}
	return s, nil
}
