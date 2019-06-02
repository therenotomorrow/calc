package calc

import (
	"math"
	"testing"
)

func equals(w, g float64) bool {
	return math.IsNaN(w) == math.IsNaN(g) || math.Abs(w-g) > 1e-6 // Precision
}

func TestSumPerform(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
		err  error
	}{
		{"empty", args{[]float64{}}, nan, notEnoughErr},
		{"one element", args{[]float64{1}}, nan, notEnoughErr},
		{"two element", args{[]float64{1.1, -2.2}}, -1.1, nil},
		{"smoke", args{[]float64{1.1, -2.2, 3.3, -4.4}}, -2.2, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sum{}
			got, err := s.Perform(tt.args.nums...)
			if err != tt.err {
				t.Errorf("Sum.Perform() error = %v, want %v", err, tt.err)
				return
			}
			if !equals(got, tt.want) {
				t.Errorf("Sum.Perform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubPerform(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
		err  error
	}{
		{"empty", args{[]float64{}}, nan, notEnoughErr},
		{"one element", args{[]float64{1}}, nan, notEnoughErr},
		{"two element", args{[]float64{1.1, -2.2}}, 3.3, nil},
		{"usual", args{[]float64{1.1, -2.2, 3.3, -4.4}}, 4.4, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sub{}
			got, err := s.Perform(tt.args.nums...)
			if err != tt.err {
				t.Errorf("Sub.Perform() error = %v, want %v", err, tt.err)
				return
			}
			if !equals(got, tt.want) {
				t.Errorf("Sub.Perform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulPerform(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
		err  error
	}{
		{"empty", args{[]float64{}}, nan, notEnoughErr},
		{"one element", args{[]float64{1}}, nan, notEnoughErr},
		{"two element", args{[]float64{1.1, -2}}, -2.2, nil},
		{"usual", args{[]float64{1.1, -2.2, 3.3, -4.4}}, 35.1384, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Mul{}
			got, err := s.Perform(tt.args.nums...)
			if err != tt.err {
				t.Errorf("Mul.Perform() error = %v, want %v", err, tt.err)
				return
			}
			if !equals(got, tt.want) {
				t.Errorf("Mul.Perform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivPerform(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
		err  error
	}{
		{"empty", args{[]float64{}}, nan, notEnoughErr},
		{"one element", args{[]float64{1}}, nan, notEnoughErr},
		{"two element", args{[]float64{1.1, -2}}, -0.55, nil},
		{"usual", args{[]float64{1.1, -2.2, 3.2, -4}}, 0.0390625, nil},
		// special cases
		{"zero element in middle", args{[]float64{1.1, 0, 3.3}}, nan, divByZeroErr},
		{"zero element at first", args{[]float64{0.0, -2.2, -4.4}}, 0, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Div{}
			got, err := s.Perform(tt.args.nums...)
			if err != tt.err {
				t.Errorf("Div.Perform() error = %v, want %v", err, tt.err)
				return
			}
			if !equals(got, tt.want) {
				t.Errorf("Div.Perform() = %v, want %v", got, tt.want)
			}
		})
	}
}
