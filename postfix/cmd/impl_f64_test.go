package cmd

import (
	"fmt"
	"math"
	"testing"
)

func TestParseF64(t *testing.T) {
	var appendSI = func(s string, si rune) string {
		return fmt.Sprintf("%s%c", s, si)
	}
	datas := []struct {
		in   string
		want float64
	}{
		{"0.0", 0.0},
		{"0", 0.0},
		{"1_000.0", 1_000.0},
		{"1,000.0", 1_000.0},
		{"-100.0", -100.0},
		{Pi, math.Pi},
		{appendSI("1", Kilo), 1_000.0},
		{appendSI("1", Mega), 1_000_000.0},
		{appendSI("1", Giga), 1_000_000_000.0},
		{appendSI("1", Milli), 0.001},
		{appendSI("1", Micro), 0.00_000_1},
		{appendSI("1", Nano), 0.00_000_000_1},
		{appendSI("1_000", Kilo), 1_000_000.0},
		{appendSI("-1_000", Kilo), -1_000_000.0},
	}

	for _, data := range datas {
		got, err := parseF64(data.in)
		if err != nil {
			t.Fatalf("err:%v, in:%v", err, data.in)
		}
		if got != data.want {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
	}

	invalidDatas := []struct {
		in         string
		wantErrMsg string
	}{
		{"a", newErrInvalidId("a").Error()},
		{appendSI("a", Kilo), newErrInvalidId("ak").Error()},
	}

	for _, data := range invalidDatas {
		got, err := parseF64(data.in)
		if err == nil {
			t.Fatalf("in:%v, wantErrMsg:%v, got:%v", data.in, data.wantErrMsg, got)
		}
		if err.Error() != data.wantErrMsg {
			t.Fatalf("in:%v, wantErrMsg:%v, err:%v", data.in, data.wantErrMsg, err)
		}
	}
}
