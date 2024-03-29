package datatypes

type Pair[T, K any] struct {
	First  T
	Second K
}

type Triple[E, R, T any] struct {
	Left   E
	Middle R
	Right  T
}
