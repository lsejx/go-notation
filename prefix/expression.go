package prefix

type expr[T any] interface {
	value() (T, error)
}

type value[T any] struct {
	v T
}

func newValue[T any](v T) *value[T] {
	return &value[T]{
		v: v,
	}
}

func (v *value[T]) value() (T, error) {
	return v.v, nil
}

type operation[T any] struct {
	f      func(T, T) (T, error)
	parent *operation[T]
	l      expr[T]
	r      expr[T]
}

func newOperation[T any](f func(T, T) (T, error)) *operation[T] {
	return &operation[T]{
		f:      f,
		parent: nil,
		l:      nil,
		r:      nil,
	}
}

func (o *operation[T]) setL(l *operation[T]) {
	if o.l != nil {
		panic("development error: o.l != nil")
	}
	o.l = l
	l.parent = o
}

func (o *operation[T]) setR(r *operation[T]) {
	if o.r != nil {
		panic("development error: o.r != nil")
	}
	o.r = r
	r.parent = o
}

func (o *operation[T]) value() (T, error) {
	if o.f == nil {
		panic("development error: o.f == nil")
	}
	if o.l == nil || o.r == nil {
		return defV[T](), ErrTooFewOperand
	}
	lv, err := o.l.value()
	if err != nil {
		return defV[T](), err
	}
	rv, err := o.r.value()
	if err != nil {
		return defV[T](), err
	}
	return o.f(lv, rv)
}
