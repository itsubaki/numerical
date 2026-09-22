package numerical

// Gradient returns the numerical gradient of f at x.
func Gradient[T float32 | float64](
	f func(x ...T) T,
	x []T,
	h T,
) []T {
	xh := append([]T(nil), x...)

	grads := make([]T, len(x))
	for i := range x {
		xi := x[i]

		// fxh1
		xh[i] = xi + h
		fxh1 := f(xh...)

		// fxh2
		xh[i] = xi - h
		fxh2 := f(xh...)

		// grads
		grads[i] = (fxh1 - fxh2) / (2 * h)

		// revert
		xh[i] = xi
	}

	return grads
}
