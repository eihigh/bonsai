package iter

type Yield[T any] func(T) bool

type Seq[T any] func(Yield[T]) bool

// Append appends the values from seq to the slice and returns the extended slice.
// TODO: Replace with slices.Append from Go 1.23
func Append[Slice ~[]Elem, Elem any](x Slice, seq Seq[Elem]) Slice {
	seq(func(v Elem) bool {
		x = append(x, v)
		return true
	})
	return x
}

// Collect collects values from seq into a new slice and returns it.
// TODO: Replace with slices.Collect from Go 1.23
func Collect[Elem any](seq Seq[Elem]) []Elem {
	return Append([]Elem(nil), seq)
}

func (y Yield[T]) Skip(n int) bool {
	var zero T
	for range n {
		if !y(zero) {
			return false
		}
	}
	return true
}

func (y Yield[T]) Repeat(n int, v T) bool {
	for range n {
		if !y(v) {
			return false
		}
	}
	return true
}
