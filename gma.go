package gma

type AssertFunc[T any] func(x T) bool

type MatchFunc[T, P any] func(want P, got T) bool

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

type Match[T, P any] struct {
	match MatchFunc[T, P]
	want  P
	msg   string
}

func MATCH[T, P any](matcher MatchFunc[T, P], want P, msg ...string) *Match[T, P] {
	m := &Match[T, P]{
		match: matcher,
		want:  want,
	}
	if len(msg) > 0 {
		m.msg = msg[0]
	}
	return m
}

func (m *Match[T, P]) Matches(x any) bool {
	return m.match(m.want, x.(T))
}

func (m *Match[T, P]) String() string {
	if m.msg != "" {
		return m.msg
	}

	return "matcher"
}
