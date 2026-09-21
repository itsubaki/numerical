package numerical_test

import (
	"fmt"

	"github.com/itsubaki/numerical"
)

func ExampleGradient() {
	f := func(x ...float64) float64 {
		return x[0]*x[0] + x[1]*x[1]
	}

	for _, v := range numerical.Gradient(f, []float64{
		1, 2,
	}, 1e-8) {
		fmt.Printf("%.2f\n", v)
	}

	// Output:
	// 2.00
	// 4.00
}
