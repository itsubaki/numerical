package numerical_test

import (
	"fmt"

	"github.com/itsubaki/numerical"
)

func ExampleRK4() {
	// dy/dx = x + y, y(0) = 1
	// y = 2*exp(x) - x - 1
	dydx := func(x, y float64) float64 {
		return x + y
	}

	y := numerical.RK4(dydx, 0, 1, 0.1)
	fmt.Printf("%.6f\n", y)

	// Output:
	// 1.110342
}
