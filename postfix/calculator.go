package postfix

import (
	"errors"
)

var (
	ErrTooFewOperand  error = errors.New("operand is not enough")
	ErrTooFewOperator error = errors.New("operator is not enough")
	ErrNoFormula      error = errors.New("no formula")
)

func getDefaultVal[T any]() T {
	var v T
	return v
}

func pop[T any](slice *[]T) (T, error) {
	if len(*slice) == 0 {
		return getDefaultVal[T](), errors.New("length of the slice is 0")
	}
	v := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]

	return v, nil
}

type Calculator[T any] struct {
	operandStack []T
}

// NewCalculator returns Calculator with specified capacity.
// Internal stack size may increase because of builtin append function.
func NewCalculator[T any](capacity int) *Calculator[T] {
	return &Calculator[T]{make([]T, 0, capacity)}
}

// AppendValue appends a value to internal stack by calling builtin append function.
func (p *Calculator[T]) AppendValue(value T) {
	p.operandStack = append(p.operandStack, value)
}

type operandsOfAOperation[T any] struct {
	left  T
	right T
}

func (p *Calculator[T]) popTwo() (operandsOfAOperation[T], error) {
	right, err := pop(&p.operandStack)
	if err != nil {
		return operandsOfAOperation[T]{}, ErrTooFewOperand
	}

	left, err := pop(&p.operandStack)
	if err != nil {
		return operandsOfAOperation[T]{}, ErrTooFewOperand
	}

	return operandsOfAOperation[T]{left, right}, nil
}

// Operate pops two values from internal stack and pass them into operation.
// Non-nil error when stack has less than two values or operation returned non-nil error.
// This panics if operation == nil.
func (p *Calculator[T]) Operate(operation func(T, T) (T, error)) (T, error) {
	if operation == nil {
		panic("github.com/lsejx/go-notation/postfix: operation == nil")
	}
	operands, err := p.popTwo()
	if err != nil {
		return getDefaultVal[T](), err
	}

	operationResult, err := operation(operands.left, operands.right)
	if err != nil {
		return getDefaultVal[T](), err
	}

	p.operandStack = append(p.operandStack, operationResult)

	return operationResult, nil
}

// Result returns value in internal stack.
// This function is intended to be used at the end of the formula.
// Non-nil error when length of the stack == 0 or > 1.
// This will not clears the stack, so the returned value will keep being in the stack.
func (p *Calculator[T]) Result() (T, error) {
	if len(p.operandStack) == 0 {
		return getDefaultVal[T](), ErrNoFormula
	}
	if len(p.operandStack) > 1 {
		return getDefaultVal[T](), ErrTooFewOperator
	}
	return p.operandStack[0], nil
}

// Clear allocates new stack and set it into the Calculator.
func (p *Calculator[T]) Clear(newCapacity int) {
	p.operandStack = make([]T, newCapacity)
}
