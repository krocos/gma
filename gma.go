package gma

type AssertFunc[T any] func(x T) bool

type MatchFunc[T any] func(want T, got T) bool

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

func (m *Assert[T]) Matches(x any) bool {
	return m.asserter(x.(T))
}

func (m *Assert[T]) String() string {
	if m.msg != "" {
		return m.msg
	}

	return "asserter"
}

type Match[T any] struct {
	match MatchFunc[T]
	want  T
	msg   string
}

func MATCH[T any](matcher MatchFunc[T], want T, msg ...string) *Match[T] {
	m := &Match[T]{
		match: matcher,
		want:  want,
	}
	if len(msg) > 0 {
		m.msg = msg[0]
	}
	return m
}

func (m *Match[T]) Matches(x any) bool {
	return m.match(m.want, x.(T))
}

func (m *Match[T]) String() string {
	if m.msg != "" {
		return m.msg
	}

	return "matcher"
}
