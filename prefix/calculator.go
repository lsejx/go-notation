package prefix

import "errors"

var (
	ErrTooFewOperator = errors.New("operator is not enough")
	ErrTooFewOperand  = errors.New("operand is not enough")
	ErrNoFormula      = errors.New("no formula")
)

func defV[T any]() T {
	var v T
	return v
}

type Calculator[T any] struct {
	root   expr[T]
	cursor *operation[T]
}

// NewCalculator returns new empty Calculator.
func NewCalculator[T any]() *Calculator[T] {
	return &Calculator[T]{
		root:   nil,
		cursor: nil,
	}
}

// AppendValue sets the value to latest operation.
// If no operation is inputted before, the value stands as result of the calculation.
func (c *Calculator[T]) AppendValue(value T) error {
	val := newValue(value)
	if c.root == nil {
		c.root = val
		return nil
	}
	if c.cursor == nil {
		return ErrTooFewOperator
	}
	if c.cursor.l == nil {
		c.cursor.l = val
		return nil
	}
	if c.cursor.r == nil {
		c.cursor.r = val
		c.cursor = c.cursor.parent
		return nil
	}
	c.cursor = c.cursor.parent
	return c.AppendValue(value)
}

// AppendValue sets the operation to latest operation as an operand.
// If no operation is inputted before, the operation's result stands as result of the calculation.
// This panics if f == nil.
func (c *Calculator[T]) AppendOperation(f func(lhs, rhs T) (T, error)) error {
	if f == nil {
		panic("github.com/lsejx/go-notation/prefix: f == nil")
	}
	o := newOperation(f)
	if c.root == nil {
		c.root = o
		c.cursor = o
		return nil
	}
	if c.cursor == nil {
		return ErrTooFewOperator
	}
	if c.cursor.l == nil {
		c.cursor.setL(o)
		c.cursor = o
		return nil
	}
	if c.cursor.r == nil {
		c.cursor.setR(o)
		c.cursor = o
		return nil
	}
	c.cursor = c.cursor.parent
	return c.AppendOperation(f)
}

// Run returns result of calculation.
// If the Calculator is empty, ErrNoFormula is returned.
func (c *Calculator[T]) Run() (T, error) {
	if c.root == nil {
		return defV[T](), ErrNoFormula
	}
	return c.root.value()
}
