package calc

import (
	"errors"
	"math"
)

// Op represents calculation of the elements of the input numbers.
type Op interface {
	Perform(...float64) (float64, error)
}

// Sum type represents the sum operation.
type Sum struct{}

func (*Sum) Perform(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return nan, notEnoughErr
	}
	r := nums[0]
	for _, a := range nums[1:] {
		r += a
	}
	return r, nil
}

// Sub type represents the subtraction operation.
type Sub struct{}

func (*Sub) Perform(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return nan, notEnoughErr
	}
	r := nums[0]
	for _, a := range nums[1:] {
		r -= a
	}
	return r, nil
}

// Mul type represents the multiplication operation.
type Mul struct{}

func (*Mul) Perform(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return nan, notEnoughErr
	}
	r := nums[0]
	for _, a := range nums[1:] {
		r *= a
	}
	return r, nil
}

// Div type represents the division operation.
type Div struct{}

func (*Div) Perform(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return nan, notEnoughErr
	}
	r := nums[0]
	for _, a := range nums[1:] {
		if a == 0 {
			return nan, divByZeroErr
		}
		r /= a
	}
	// fix -0 of IEEE754
	if r == 0 {
		r = math.Abs(r)
	}
	return r, nil
}

var (
	notEnoughErr = errors.New("not enough values, must be 2 or more")
	divByZeroErr = errors.New("division by zero")

	// shortcut for nan value
	nan = math.NaN()
)
