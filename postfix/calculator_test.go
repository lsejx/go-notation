package postfix

import (
	"bytes"
	"errors"
	"testing"
)

func TestPop(t *testing.T) {
	datas := []struct {
		in    []byte
		after []byte
		want  byte
	}{
		{[]byte{1, 2, 3, 4, 5}, []byte{1, 2, 3, 4}, 5},
	}

	for _, data := range datas {
		got, err := pop(&data.in)
		if err != nil {
			t.Fatalf("err:%v, in:%v", err, data.in)
		}
		if data.want != got {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
		if !bytes.Equal(data.after, data.in) {
			t.Fatalf("after:%v, actual:%v", data.after, data.in)
		}
	}
}

func TestNewCalculator(t *testing.T) {
	datas := []struct {
		capacity int
	}{
		{0},
		{1},
		{2},
		{3},
	}

	for _, data := range datas {
		c := NewCalculator[any](data.capacity)
		actual := cap(c.operandStack)
		if data.capacity != actual {
			t.Fatalf("capacity:%v, actual:%v", data.capacity, actual)
		}
	}
}

func TestAppendOperand(t *testing.T) {
	datas := []struct {
		ins   []byte
		after []byte
	}{
		{[]byte{1}, []byte{1}},
		{[]byte{1, 2}, []byte{1, 2}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3}},
		{[]byte{1, 2, 3, 4}, []byte{1, 2, 3, 4}},
	}

	for _, data := range datas {
		c := NewCalculator[byte](len(data.ins))
		for _, in := range data.ins {
			c.AppendOperand(in)
		}
		if !bytes.Equal(data.after, c.operandStack) {
			t.Fatalf("after:%v, actual:%v", data.after, c.operandStack)
		}
	}
}

func TestPopTwo(t *testing.T) {
	datas := []struct {
		stack []byte
		want  operandsOfAOperation[byte]
	}{
		{[]byte{1, 2}, operandsOfAOperation[byte]{1, 2}},
		{[]byte{1, 2, 3}, operandsOfAOperation[byte]{2, 3}},
	}

	for _, data := range datas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.popTwo()
		if err != nil {
			t.Fatalf("err:%v, stack:%v, operands:%v", err, data.stack, c.operandStack)
		}
		if got != data.want {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
	}

	invalidDatas := []struct {
		stack   []byte
		wantErr error
	}{
		{[]byte{}, ErrTooFewOperands},
		{[]byte{1}, ErrTooFewOperands},
	}

	for _, data := range invalidDatas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.popTwo()
		if err == nil {
			t.Fatalf("stack:%v, got:%v", data.stack, got)
		}
		if data.wantErr != err {
			t.Fatalf("wantErr:%v, err:%v", data.wantErr, err)
		}
	}
}

func TestOperate(t *testing.T) {
	datas := []struct {
		stack     []byte
		operation func(byte, byte) (byte, error)
		want      byte
		after     []byte
	}{
		{[]byte{1, 2}, func(b1, b2 byte) (byte, error) { return b1 + b2, nil }, 3, []byte{3}},
		{[]byte{1, 2, 3}, func(b1, b2 byte) (byte, error) { return b1 + b2, nil }, 5, []byte{1, 5}},
	}

	for _, data := range datas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.Operate(data.operation)
		if err != nil {
			t.Fatalf("err:%v, data:%v", err, data)
		}
		if data.want != got {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
		if !bytes.Equal(data.after, c.operandStack) {
			t.Fatalf("after:%v, actual:%v", data.after, c.operandStack)
		}
	}

	testErr := errors.New("test")
	invalidDatas := []struct {
		stack     []byte
		operation func(byte, byte) (byte, error)
		wantErr   error
	}{
		{[]byte{1}, func(byte, byte) (byte, error) { return 0, nil }, ErrTooFewOperands},
		{[]byte{1, 2}, func(b1, b2 byte) (byte, error) { return 0, testErr }, testErr},
	}

	for _, data := range invalidDatas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.Operate(data.operation)
		if err == nil {
			t.Fatalf("got:%v, data:%v", got, data)
		}
		if err != data.wantErr {
			t.Fatalf("err:%v, wantErr:%v", err, data.wantErr)
		}
	}
}

func TestResult(t *testing.T) {
	datas := []struct {
		stack []byte
		want  byte
	}{
		{[]byte{1}, 1},
		{[]byte{2}, 2},
	}

	for _, data := range datas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.Result()
		if err != nil {
			t.Fatalf("err:%v, stack:%v", err, data.stack)
		}
		if data.want != got {
			t.Fatalf("got:%v, data:%v", got, data)
		}
	}

	invalidDatas := []struct {
		stack   []byte
		wantErr error
	}{
		{[]byte{}, ErrNoFormula},
		{[]byte{1, 2}, ErrTooFewOperators},
	}

	for _, data := range invalidDatas {
		c := &Calculator[byte]{bytes.Clone(data.stack)}
		got, err := c.Result()
		if err == nil {
			t.Fatalf("got:%v, stack:%v", got, data.stack)
		}
		if data.wantErr != err {
			t.Fatalf("wantErr:%v, err:%v", data.wantErr, err)
		}
	}
}

func TestClear(t *testing.T) {
	c := &Calculator[byte]{[]byte{1, 2, 3, 4, 5}}
	c.Clear(0)
	if !bytes.Equal(c.operandStack, []byte{}) {
		t.Fatalf("operandStack:%v", c.operandStack)
	}
}
