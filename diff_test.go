package numerical_test

import (
	"fmt"

	"github.com/itsubaki/numerical"
)

func ExampleDiff() {
	f := func(x float64) float64 {
		return x * x
	}

	for _, v := range numerical.Diff(f, []float64{
		1, 2, 3, 4,
	}, 1e-8) {
		fmt.Printf("%.2f\n", v)
	}

	// Output:
	// 2.00
	// 4.00
	// 6.00
	// 8.00
}
