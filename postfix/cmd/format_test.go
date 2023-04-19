package cmd

import "testing"

func TestInsertComma(t *testing.T) {
	datas := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"1", "1"},
		{"10", "10"},
		{"100", "100"},
		{"1000", "1,000"},
		{"10000", "10,000"},
		{"100000", "100,000"},
		{"1000000", "1,000,000"},
	}

	for _, data := range datas {
		got := insertComma(data.in)
		if got != data.want {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
	}
}

func TestFormat(t *testing.T) {
	datas := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"10", "10"},
		{"1000", "1,000"},
		{"10000", "10,000"},
		{"1000000", "1,000,000"},
		{"-10", "-10"},
		{"-1000", "-1,000"},
		{"-10000", "-10,000"},
		{"-1000000", "-1,000,000"},
		{"1.01", "1.01"},
		{"1.0001", "1.0001"},
		{"-1.01", "-1.01"},
		{"-1.0001", "-1.0001"},
		{"1000.0001", "1,000.0001"},
		{"-1000.0001", "-1,000.0001"},
	}

	for _, data := range datas {
		got, err := format(data.in)
		if err != nil {
			t.Fatalf("err:%v, in:%v", err, data.in)
		}
		if got != data.want {
			t.Fatalf("want:%v, got:%v", data.want, got)
		}
	}

	invalidDatas := []struct {
		in      string
		wantMsg string
	}{
		{"+", "only a sign, there is not number"},
		{"-", "only a sign, there is not number"},
		{"1234a", "a is invalid"},
		{"1,000", ", is invalid"},
		{"1.000.0", "there are two periods"},
	}

	for _, data := range invalidDatas {
		got, err := format(data.in)
		if err == nil {
			t.Fatalf("in:%v, wantMsg:%v, got:%v", data.in, data.wantMsg, got)
		}
		if err.Error() != data.wantMsg {
			t.Fatalf("in:%v, wantMsg:%v, errMsg:%v", data.in, data.wantMsg, err.Error())
		}
	}
}
