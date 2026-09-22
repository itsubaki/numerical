package numerical

func Diff[T float32 | float64](
	f func(x T) T,
	x []T,
	h T,
) []T {
	diff := make([]T, len(x))
	for i, xi := range x {
		diff[i] = (f(xi+h) - f(xi-h)) / (2 * h)
	}

	return diff
}
