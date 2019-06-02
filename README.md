calc
====

Calculator based on [Pipelines](https://en.wikipedia.org/wiki/Pipeline_\(computing\)).

---

Usage
-----

```go
package main

import (
	"fmt"
	"github.com/kxnes/calc"
	"os"
	"strings"
)

func main() {
	c := calc.New(map[rune]calc.Op{
		'+': &calc.Sum{},
		'-': &calc.Sub{},
		'*': &calc.Mul{},
		'/': &calc.Div{},
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
```

Extend
------

For extending, just implement `calc.Op` interface like example below.

```go
// Sin type represents the sin operation.
type Sin struct{}

func (*Sin) Perform(nums ...float64) (float64, error) {
	if len(nums) != 1 {
		return nan, errors.New("operation \"Sin\" may perform on 1 number")
	}
	return 	math.Sin(nums[0]), nil
}
```
