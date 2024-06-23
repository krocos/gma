package gma

type AssertFunc[T any] func(x T) bool

type Assert[T any] struct {
	asserter AssertFunc[T]
	msg      string
}

func ASSERT[T any](asserter AssertFunc[T], msg ...string) *Assert[T] {
	m := &Assert[T]{asserter: asserter}
	if len(msg) > 0 {
		m.msg = msg[0]
	}
	return m
}

func (a *Assert[T]) Matches(x any) bool {
	return a.asserter(x.(T))
}

func (a *Assert[T]) String() string {
	return a.msg
}
