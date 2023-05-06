package fp

func Unfold[T any](init T, f func(T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			out <- init
			init = f(init)
		}
	}()

	return out
}

func Unfold2[T any](u0, u1 T, f func(T, T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		out <- u0
		out <- u1

		for {
			tmp := f(u0, u1)
			out <- tmp
			u0 = u1
			u1 = tmp
		}
	}()

	return out
}

func Fold[I, O any](in <-chan I, init O, f func(I, O) O) O {
	out := init

	for v := range in {
		out = f(v, out)
	}

	return out
}

func TakeWhile[T any](in <-chan T, f func(T) bool) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			if f(v) {
				out <- v
			} else {
				break
			}
		}
	}()

	return out
}

func Map[I, O any](in <-chan I, f func(I) O) chan O {
	out := make(chan O)

	go func() {
		defer close(out)

		for v := range in {
			out <- f(v)
		}
	}()

	return out
}

func Filter[T any](in <-chan T, f func(T) bool) chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for v := range in {
			if f(v) {
				out <- v
			}
		}
	}()

	return out
}
