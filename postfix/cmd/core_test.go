package cmd

import (
	"errors"
	"testing"

	"github.com/lsejx/go-notation/postfix"
)

func TestNewErrInvalidId(t *testing.T) {
	datas := []struct {
		in   []string
		want string
	}{
		{[]string{"a"}, "invalid: a"},
		{[]string{"a", "b"}, "invalid: a b"},
	}

	for _, data := range datas {
		err := newErrInvalidId(data.in...)
		if err.Error() != data.want {
			t.Fatalf("in:%v, want:%v, got:%v", data.in, data.want, err)
		}
	}
}

func TestRunner(t *testing.T) {
	errTest := errors.New("test")
	operations := map[string]func(int, int) (int, error){
		"success": func(i1, i2 int) (int, error) {
			return 1, nil
		},
		"error": func(i1, i2 int) (int, error) {
			return 0, errTest
		},
	}
	parse := func(string) (int, error) {
		return 2, nil
	}
	format := func(int) (string, error) {
		return "abc", nil
	}

	datas := []struct {
		formula []string
		want    string
		wantErr error
	}{
		{[]string{"a", "a", "success"}, "abc", nil},
		{[]string{"a", "a", "error"}, "", errTest},
		{[]string{"a", "success"}, "", postfix.ErrTooFewOperand},
		{[]string{"a", "a"}, "", postfix.ErrTooFewOperator},
	}

	for _, data := range datas {
		runner := newRunner(operations, parse, format)
		got, err := runner.Run(data.formula)
		if !errors.Is(err, data.wantErr) {
			t.Fatalf("err:%v, wantErr:%v, formula:%v", err, data.wantErr, data.formula)
		}
		if got != data.want {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
	}
}
