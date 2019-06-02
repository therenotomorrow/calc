package calc

import (
	"errors"
	"math"
)

// Calc represents interface for calculation
type Calc interface {
	R() float64              // R returns result of the calculation
	E() error                // E returns error of the calculation
	C(rune, ...float64) Calc // C applies the operation to numbers
}

// internal implementation of Calc interface
type impl struct {
	ops map[rune]Op
	res float64
	err error
}

func (c *impl) R() float64 {
	r := c.res

	if c.err != nil {
		r = nan
	}

	c.res = nan
	return r
}

func (c *impl) E() error {
	return c.err
}

func (c *impl) C(opname rune, nums ...float64) Calc {

	if c.err != nil {
		c.res = nan
		return c
	}

	if !math.IsNaN(c.res) {
		nums = append([]float64{c.res}, nums...)
	}

	if op, ok := c.ops[opname]; !ok {
		c.res, c.err = nan, OperationNotExistErr
	} else {
		c.res, c.err = op.Perform(nums...)
	}

	return c
}

func New(ops map[rune]Op) Calc {
	return &impl{ops: ops, res: nan}
}

var OperationNotExistErr = errors.New("operation is not exist")
