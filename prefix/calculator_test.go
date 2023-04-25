package prefix

import (
	"errors"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	c := NewCalculator[int]()
	if c.root != nil {
		t.Fatal("nonnil root")
	}
	if c.cursor != nil {
		t.Fatal("nonnil cursor")
	}
}

func TestAppendValue(t *testing.T) {
	o1 := &operation[int]{nil, nil, nil, nil}
	o2 := &operation[int]{nil, nil, &value[int]{3}, nil}
	o3 := &operation[int]{nil, nil, nil, nil}
	o4 := &operation[int]{nil, o3, &value[int]{1}, &value[int]{2}}
	o3.l = o4
	datas := []struct {
		id   string
		in   int
		c    *Calculator[int]
		want func(*Calculator[int]) int
	}{
		{
			"nil calculator",
			-100,
			&Calculator[int]{nil, nil},
			func(c *Calculator[int]) int {
				if c.cursor != nil {
					t.FailNow()
				}
				return c.root.(*value[int]).v
			},
		},
		{
			"operation left",
			-101,
			&Calculator[int]{o1, o1},
			func(c *Calculator[int]) int {
				if c.cursor == nil {
					t.FailNow()
				}
				return o1.l.(*value[int]).v
			},
		},
		{
			"operation right",
			-102,
			&Calculator[int]{o2, o2},
			func(c *Calculator[int]) int {
				if c.cursor != nil {
					t.FailNow()
				}
				return o2.r.(*value[int]).v
			},
		},
		{
			"recursive",
			-103,
			&Calculator[int]{o3, o4},
			func(c *Calculator[int]) int {
				if c.cursor != nil {
					t.FailNow()
				}
				return o3.r.(*value[int]).v
			},
		},
	}

	for _, data := range datas {
		err := data.c.AppendValue(data.in)
		if err != nil {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
		if data.in != data.want(data.c) {
			t.Fatalf("in:%v, want:%v", data.in, data.want(data.c))
		}
	}

	o5 := &operation[int]{nil, nil, &value[int]{20}, &value[int]{23}}
	invalidDatas := []struct {
		id      string
		in      int
		c       *Calculator[int]
		wantErr error
	}{
		{
			"root value",
			100,
			&Calculator[int]{&value[int]{50}, nil},
			ErrTooFewOperator,
		},
		{
			"full",
			105,
			&Calculator[int]{o5, o5},
			ErrTooFewOperator,
		},
	}

	for _, data := range invalidDatas {
		err := data.c.AppendValue(data.in)
		if err == nil {
			t.Fatalf("nilerr, id:%v", data.id)
		}
		if !errors.Is(err, data.wantErr) {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
	}
}

func TestAppendOperation(t *testing.T) {
	o1 := &operation[int]{nil, nil, nil, nil}
	o2 := &operation[int]{nil, nil, &value[int]{10}, nil}
	o3 := &operation[int]{nil, nil, nil, nil}
	o4 := &operation[int]{nil, o3, &value[int]{11}, &value[int]{12}}
	o3.l = o4
	datas := []struct {
		id     string
		in     func(int, int) (int, error)
		c      *Calculator[int]
		assert func(*Calculator[int])
	}{
		{
			"nil",
			func(i1, i2 int) (int, error) { return 101, nil },
			&Calculator[int]{nil, nil},
			func(c *Calculator[int]) {
				if c.root.(*operation[int]) != c.cursor {
					t.Fatal("nil")
				}
				v, e := c.cursor.f(0, 0)
				if e != nil {
					t.FailNow()
				}
				if v != 101 {
					t.FailNow()
				}
			},
		},
		{
			"left",
			func(i1, i2 int) (int, error) { return 102, nil },
			&Calculator[int]{o1, o1},
			func(c *Calculator[int]) {
				if c.root.(*operation[int]).l == nil {
					t.Fatal("left")
				}
				v, e := c.cursor.f(0, 0)
				if e != nil {
					t.FailNow()
				}
				if v != 102 {
					t.FailNow()
				}
			},
		},
		{
			"right",
			func(i1, i2 int) (int, error) { return 103, nil },
			&Calculator[int]{o2, o2},
			func(c *Calculator[int]) {
				if c.root.(*operation[int]).r == nil {
					t.Fatal("right")
				}
				v, e := c.cursor.f(0, 0)
				if e != nil {
					t.FailNow()
				}
				if v != 103 {
					t.FailNow()
				}
			},
		},
		{
			"recursive",
			func(i1, i2 int) (int, error) { return 105, nil },
			&Calculator[int]{o3, o4},
			func(c *Calculator[int]) {
				if c.cursor != c.root.(*operation[int]).r {
					t.Fatal("recursive")
				}
				v, e := c.cursor.f(0, 0)
				if e != nil {
					t.FailNow()
				}
				if v != 105 {
					t.FailNow()
				}
			},
		},
	}

	for _, data := range datas {
		err := data.c.AppendOperation(data.in)
		if err != nil {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
		data.assert(data.c)
	}

	o5 := &operation[int]{nil, nil, &value[int]{5}, &value[int]{7}}
	invalidDatas := []struct {
		id      string
		in      func(int, int) (int, error)
		c       *Calculator[int]
		wantErr error
	}{
		{
			"root value",
			func(i1, i2 int) (int, error) { return 1, nil },
			&Calculator[int]{&value[int]{1}, nil},
			ErrTooFewOperator,
		},
		{
			"full",
			func(i1, i2 int) (int, error) { return 2, nil },
			&Calculator[int]{o5, o5},
			ErrTooFewOperator,
		},
	}

	for _, data := range invalidDatas {
		err := data.c.AppendOperation(data.in)
		if err == nil {
			t.Fatalf("nilerr, id:%v", data.id)
		}
		if !errors.Is(err, data.wantErr) {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}

	}
}

func TestRun(t *testing.T) {
	o1 := &operation[int]{func(i1, i2 int) (int, error) { return i1 - i2, nil }, nil, &value[int]{1}, &value[int]{2}}
	datas := []struct {
		id   string
		c    *Calculator[int]
		want int
	}{
		{
			"value",
			&Calculator[int]{&value[int]{-999}, nil},
			-999,
		},
		{
			"operation",
			&Calculator[int]{o1, o1},
			-1,
		},
	}

	for _, data := range datas {
		res, err := data.c.Run()
		if err != nil {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
		if res != data.want {
			t.Fatalf("res:%v, id:%v", res, data.id)
		}
	}

	invalidDatas := []struct {
		id      string
		c       *Calculator[int]
		wantErr error
	}{
		{
			"nil",
			&Calculator[int]{nil, nil},
			ErrNoFormula,
		},
	}

	for _, data := range invalidDatas {
		res, err := data.c.Run()
		if err == nil {
			t.Fatalf("nilerr, res:%v, id:%v", res, data.id)
		}
		if !errors.Is(err, data.wantErr) {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
	}
}

func TestCalculator(t *testing.T) {
	fatal := func(e error) {
		if e != nil {
			t.Fatal(e)
		}
	}
	datas := []struct {
		id   string
		f    func(*Calculator[int])
		want int
	}{
		{
			"one value",
			func(c *Calculator[int]) {
				fatal(c.AppendValue(100))
			},
			100,
		},
		{
			"1-2",
			func(c *Calculator[int]) {
				fatal(c.AppendOperation(func(lhs, rhs int) (int, error) { return lhs - rhs, nil }))
				fatal(c.AppendValue(1))
				fatal(c.AppendValue(2))
			},
			-1,
		},
		{
			"1-2-3",
			func(c *Calculator[int]) {
				sub := func(i1, i2 int) (int, error) { return i1 - i2, nil }
				fatal(c.AppendOperation(sub))
				fatal(c.AppendOperation(sub))
				fatal(c.AppendValue(1))
				fatal(c.AppendValue(2))
				fatal(c.AppendValue(3))
			},
			-4,
		},
		{
			"(1-3)-(5-4)",
			func(c *Calculator[int]) {
				sub := func(i1, i2 int) (int, error) { return i1 - i2, nil }
				fatal(c.AppendOperation(sub))
				fatal(c.AppendOperation(sub))
				fatal(c.AppendValue(1))
				fatal(c.AppendValue(3))
				fatal(c.AppendOperation(sub))
				fatal(c.AppendValue(5))
				fatal(c.AppendValue(4))
			},
			-3,
		},
	}

	for _, data := range datas {
		c := NewCalculator[int]()
		data.f(c)
		res, err := c.Run()
		if err != nil {
			t.Fatalf("err:%v, id:%v", err, data.id)
		}
		if res != data.want {
			t.Fatalf("res:%v, id:%v", res, data.id)
		}
	}

}
