package prefix

import (
	"errors"
	"testing"
)

func TestNewValue(t *testing.T) {
	datas := []struct {
		in int
	}{
		{0},
		{1},
		{5},
		{100},
		{-100},
	}

	for _, data := range datas {
		v := newValue(data.in)
		if v.v != data.in {
			t.Fatalf("v:%v, in:%v", v.v, data.in)
		}
	}
}

func TestValueValue(t *testing.T) {
	datas := []struct {
		v int
	}{
		{0},
		{1},
		{5},
		{100},
		{-100},
	}

	for _, data := range datas {
		v := &value[int]{
			v: data.v,
		}
		got, err := v.value()
		if err != nil {
			t.FailNow()
		}
		if got != data.v {
			t.Fatalf("v:%v, got:%v", data.v, got)
		}
	}
}

func TestNewOperation(t *testing.T) {
	datas := []struct {
		f func(any, any) (any, error)
	}{
		{func(a1, a2 any) (any, error) { return 0, nil }},
	}

	for _, data := range datas {
		o := newOperation(data.f)
		if o.f == nil {
			t.FailNow()
		}
		if o.l != nil {
			t.FailNow()
		}
		if o.r != nil {
			t.FailNow()
		}
	}
}

func TestSetL(t *testing.T) {
	p := &operation[any]{
		f:      nil,
		parent: nil,
		l:      nil,
		r:      nil,
	}
	l := &operation[any]{
		f:      nil,
		parent: nil,
		l:      nil,
		r:      nil,
	}
	p.setL(l)
	if p.l != l {
		t.FailNow()
	}
	if l.parent != p {
		t.FailNow()
	}
}

func TestSetR(t *testing.T) {
	p := &operation[any]{
		f:      nil,
		parent: nil,
		l:      nil,
		r:      nil,
	}
	r := &operation[any]{
		f:      nil,
		parent: nil,
		l:      nil,
		r:      nil,
	}
	p.setR(r)
	if p.r != r {
		t.FailNow()
	}
	if r.parent != p {
		t.FailNow()
	}
}

func TestOperationValue(t *testing.T) {
	datas := []struct {
		id   string
		o    operation[int]
		want int
	}{
		{
			id: "just return",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 1, nil },
				l: &value[int]{0},
				r: &value[int]{0},
			},
			want: 1,
		},
		{
			id: "1 - 2",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
				l: &value[int]{1},
				r: &value[int]{2},
			},
			want: -1,
		},
		{
			id: "1 - 2 - 3",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
				l: &operation[int]{
					f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
					l: &value[int]{1},
					r: &value[int]{2},
				},
				r: &value[int]{3},
			},
			want: -4,
		},
		{
			id: "1 - 2 - 3 - 4",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
				l: &operation[int]{
					f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
					l: &operation[int]{
						f: func(i1, i2 int) (int, error) { return i1 - i2, nil },
						l: &value[int]{1},
						r: &value[int]{2},
					},
					r: &value[int]{3},
				},
				r: &value[int]{4},
			},
			want: -8,
		},
	}

	for _, data := range datas {
		res, err := data.o.value()
		if err != nil {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
		if res != data.want {
			t.Fatalf("res:%v, id:%v", res, data.id)
		}
	}

	testErr := errors.New("test")
	invalidDatas := []struct {
		id      string
		o       operation[int]
		wantErr error
	}{
		{
			id: "toofewoperand1",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, nil },
				l: nil,
				r: &value[int]{},
			},
			wantErr: ErrTooFewOperand,
		},
		{
			id: "toofewoperand2",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, nil },
				l: &value[int]{},
				r: nil,
			},
			wantErr: ErrTooFewOperand,
		},
		{
			id: "toofewoperand3",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, nil },
				l: nil,
				r: nil,
			},
			wantErr: ErrTooFewOperand,
		},
		{
			id: "lverr",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, nil },
				l: &operation[int]{
					f: func(i1, i2 int) (int, error) { return 0, testErr },
					l: &value[int]{},
					r: &value[int]{},
				},
				r: &value[int]{},
			},
			wantErr: testErr,
		},
		{
			id: "rverr",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, nil },
				l: &value[int]{},
				r: &operation[int]{
					f: func(i1, i2 int) (int, error) { return 0, testErr },
					l: &value[int]{},
					r: &value[int]{},
				},
			},
			wantErr: testErr,
		},
		{
			id: "ferr",
			o: operation[int]{
				f: func(i1, i2 int) (int, error) { return 0, testErr },
				l: &value[int]{},
				r: &value[int]{},
			},
			wantErr: testErr,
		},
	}

	for _, data := range invalidDatas {
		res, err := data.o.value()
		if !errors.Is(err, data.wantErr) {
			t.Fatalf("err:%v, wE:%v, res:%v, id:%v", err, data.wantErr, res, data.id)
		}
	}
}
