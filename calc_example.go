//+build test

package calc

import (
	"fmt"
)

func _() {
	c := New(map[rune]Op{
		'+': &Sum{},
		'-': &Sub{},
		'*': &Mul{},
		'/': &Div{},
	})

	q := c.C('*', 2, 2).C('+', 2).R()
	fmt.Printf("2 + 2 * 2 = %v\n", q)

	w := c.C('+', 2, 2).C('*', 2).R()
	fmt.Printf("(2 + 2) * 2 = %v\n", w)

	e := c.C('+', c.C('*', 2, 2).R(), c.C('*', 2, 2).R()).R()
	fmt.Printf("(2 * 2) + (2 * 2) = %v\n", e)

	r := c.C('-', c.C('+', c.C('/', 4, 2, 2).R(), c.C('*', 15, 4, 10).R()).R(), 0.01).R()
	fmt.Printf("(4 / 2 / 2) + (15 * 4 * 10) - 0.01 = %v\n", r)
}
