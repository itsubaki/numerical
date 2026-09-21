package numerical

func Diff[T float32 | float64](f func(x T) T, x []T, h T) []T {
	out := make([]T, len(x))
	for i, xi := range x {
		out[i] = (f(xi+h) - f(xi-h)) / (2 * h)
	}

	return out
}
