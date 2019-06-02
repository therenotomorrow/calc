package calc

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestImplR(t *testing.T) {
	type fields struct {
		ops map[rune]Op
		res float64
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"not error", fields{map[rune]Op{}, 1.1, nil}, 1.1},
		{"error", fields{map[rune]Op{}, 1.1, errors.New("")}, nan},
		{"not error before calculation", fields{map[rune]Op{}, nan, nil}, nan},
		{"error before calculation", fields{map[rune]Op{}, nan, errors.New("")}, nan},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &impl{
				ops: tt.fields.ops,
				res: tt.fields.res,
				err: tt.fields.err,
			}
			got := c.R()
			if !equals(got, tt.want) {
				t.Errorf("impl.R() = %v, want %v", got, tt.want)
			}
			if !math.IsNaN(c.res) {
				t.Errorf("impl.R(): res = %v, want %v", c.res, nan)
			}
		})
	}
}

func TestImplE(t *testing.T) {
	type fields struct {
		ops map[rune]Op
		res float64
		err error
	}
	errTmpl := errors.New("something went wrong")
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{"not error", fields{map[rune]Op{}, 1.1, nil}, nil},
		{"error", fields{map[rune]Op{}, 1.1, errTmpl}, errTmpl},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &impl{
				ops: tt.fields.ops,
				res: tt.fields.res,
				err: tt.fields.err,
			}
			err := c.E()
			if err != tt.err {
				t.Errorf("impl.E(): err = %v, want %v", err, tt.err)
			}
		})
	}
}

func TestImplC(t *testing.T) {
	type fields struct {
		ops map[rune]Op
		res float64
		err error
	}
	type args struct {
		opname rune
		nums   []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Calc
	}{
		{
			"operation does not exist",
			fields{map[rune]Op{'+': &Sum{}}, nan, nil},
			args{'*', []float64{1, 2, 3}},
			&impl{map[rune]Op{'+': &Sum{}}, nan, OperationNotExistErr},
		},
		{
			"operation is first in calculation",
			fields{map[rune]Op{'+': &Sum{}}, nan, nil},
			args{'+', []float64{1, 2, 3}},
			&impl{map[rune]Op{'+': &Sum{}}, 6, nil},
		},
		{
			"operation is not first in calculation",
			fields{map[rune]Op{'+': &Sum{}}, 2, nil},
			args{'+', []float64{1, 2, 3}},
			&impl{map[rune]Op{'+': &Sum{}}, 8, nil},
		},
		{
			"error occurred when operation perform",
			fields{map[rune]Op{'/': &Div{}}, 2, nil},
			args{'/', []float64{1, 0, 3}},
			&impl{map[rune]Op{'/': &Div{}}, nan, divByZeroErr},
		},
		{
			"does not process calculation when error occurred previously",
			fields{map[rune]Op{'/': &Div{}}, 2, divByZeroErr},
			args{'/', []float64{1, 0, 3}},
			&impl{map[rune]Op{'/': &Div{}}, nan, divByZeroErr},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &impl{
				ops: tt.fields.ops,
				res: tt.fields.res,
				err: tt.fields.err,
			}
			got := c.C(tt.args.opname, tt.args.nums...).(*impl)
			want := tt.want.(*impl)
			if !equals(got.res, want.res) {
				t.Errorf("impl.C(): res = %v, want %v", got.res, want.res)
			}
			if got.err != want.err {
				t.Errorf("impl.C(): err = %v, want %v", got.err, want.err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ops map[rune]Op
	}
	tests := []struct {
		name string
		args args
		want Calc
	}{
		{
			"test construct",
			args{map[rune]Op{'+': &Sum{}, '-': &Sub{}, '*': &Mul{}, '/': &Div{}}},
			&impl{map[rune]Op{'+': &Sum{}, '-': &Sub{}, '*': &Mul{}, '/': &Div{}}, nan, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.ops).(*impl)
			want := tt.want.(*impl)
			if !reflect.DeepEqual(got.ops, want.ops) {
				t.Errorf("New(): ops = %v, want %v", got.ops, want.ops)
			}
			if !equals(got.res, want.res) {
				t.Errorf("New(): res = %v, want %v", got.res, want.res)
			}
			if got.err != want.err {
				t.Errorf("New(): err = %v, want %v", got.err, want.err)
			}
		})
	}
}

func TestSmoke(t *testing.T) {
	var (
		got float64
		c   = New(map[rune]Op{
			'+': &Sum{},
			'-': &Sub{},
			'*': &Mul{},
			'/': &Div{},
		})
	)

	got = c.C('+', 2, c.C('*', 2, 2).R()).R()
	if !equals(got, 6) {
		t.Errorf("case 1: 2 + (2 * 2) = %v, want %v", got, 6)
	}

	got = c.C('+', 2, 2).C('*', 2).R()
	if !equals(got, 8) {
		t.Errorf("case 2: (2 + 2) * 2 = %v, want %v", got, 8)
	}

	got = // (2 * 2) + (2 * 2)
		c.C('+',
			// (2 * 2)
			c.C('*', 2, 2).R(),
			// (2 * 2)
			c.C('*', 2, 2).R()).R()
	if !equals(got, 8) {
		t.Errorf("case 3: (2 * 2) + (2 * 2) = %v, want %v", got, 8)
	}

	got = // (4 / 2 / 2) + (15 * 4) - 0.01
		c.C('-',
			// (4 / 2 / 2) + (15 * 4)
			c.C('+',
				// (4 / 2 / 2)
				c.C('/', 4, 2, 2).R(),
				// (15 * 4)
				c.C('*', 15, 4).R()).R(),
			0.01).R()
	if !equals(got, 60.99) {
		t.Errorf("case 4: (4 / 2 / 2) + (15 * 4) - 0.01 = %v, want %v", got, 60.99)
	}

	got = // (4 / 0 / 2) + (15 * 4) - 0.01
		c.C('-',
			// (4 / 0 / 2) + (15 * 4)
			c.C('+',
				// (4 / 0 / 2)
				c.C('/', 4, 0, 2).R(),
				// (15 * 4)
				c.C('*', 15, 4).R()).R(),
			0.01).R()
	if !math.IsNaN(got) {
		t.Errorf("case 5: (4 / 0 / 2) + (15 * 4) - 0.01 = %v, want %v", got, nan)
	}
	if c.E() == nil {
		t.Errorf("case 5: (4 / 0 / 2) + (15 * 4) - 0.01 = %v, want error", got)
	}
}
