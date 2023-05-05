package stack

// A Stack stack is a last-in-first-out (lifo) data structure.
type Stack[T any] struct {
	values []T
	size   int
}

// New returns a new stack.
func New[T any](values ...T) *Stack[T] {
	stk := Stack[T]{
		values: append(make([]T, 0, nextPow2(len(values))), values...),
		size:   len(values),
	}

	return &stk
}

// Clean frees up space that is no longer needed.
func (s *Stack[T]) Clean() *Stack[T] {
	if s.size < cap(s.values)>>1 {
		// 2*size < cap --> reduce cap to 2^n >= len for minimal n
		s.values = append(make([]T, 0, nextPow2(s.size)), s.values[:s.size]...)
	}

	return s
}

// Clear removes all values from the stack setting the size to zero.
func (s *Stack[T]) Clear() *Stack[T] {
	s.size = 0
	return s
}

// Contains determines if a value is in the stack.
func Contains[T comparable](s *Stack[T], value T) bool {
	for i := s.size - 1; 0 <= i; i-- {
		if s.values[i] == value {
			return true
		}
	}

	return false
}

// Copy returns a copy of a stack.
func (s *Stack[T]) Copy() *Stack[T] {
	cpy := Stack[T]{
		values: append(make([]T, 0, nextPow2(s.size)), s.values[:s.size]...),
		size:   s.size,
	}

	return &cpy
}

// Equal determines if two stacks are equal.
func Equal[T comparable](s0, s1 *Stack[T]) bool {
	switch {
	case s0 == s1:
		return true
	case s0.size != s1.size:
		return false
	default:
		for i := 0; i < s0.size; i++ {
			if s0.values[i] != s1.values[i] {
				return false
			}
		}

		return true
	}
}

// Peek returns the top of the stack without modifying the stack's
// state.
func (s *Stack[T]) Peek() T {
	if s.size == 0 {
		var t T
		return t
	}

	return s.values[s.size-1]
}

// Pop removes and returns the top value of the stack. If the stack is
// empty, nil will be returned.
func (s *Stack[T]) Pop() T {
	if s.size == 0 {
		var t T
		return t
	}

	s.size--

	return s.values[s.size]
}

// Push adds a value onto the stack.
func (s *Stack[T]) Push(values ...T) *Stack[T] {
	n := len(s.values) - s.size

	if 0 < n {
		if len(values) < n {
			n = len(values)
		}

		copy(s.values[s.size:], values[:n])
		s.size += n
	}

	if n < len(values) {
		s.values = append(s.values, values[n:]...)
		s.size = len(s.values)
	}

	return s
}

// Size returns the number of values on the stack.
func (s *Stack[T]) Size() int {
	return s.size
}

// Values returns a copy of the stack's values. The bottom of the stack
// is the zero index.
func (s *Stack[T]) Values() []T {
	return append(make([]T, 0, s.size), s.values[:s.size]...)
}

// --------------------------------------------------------------------
// Helpers
// --------------------------------------------------------------------

// nextPow2 returns the next power of two greater than or equal to a
// given number.
//
// 	- n < 0 --> 0
// 	- n = 0 --> 1
// 	- n > 0 --> 2^m such that 2^m >= n for  minimal m >= 0
func nextPow2(n int) int {
	// bitCap is the number of bits in an int.
	const bitCap = 32 << (^uint(0) >> 63) // Sources: bits.UintSize, strconv.IntSize

	// Source: https://web.archive.org/web/20130821015554/http://bob.allegronetwork.com/prog/tricks.html#roundtonextpowerof2

	if n == 0 {
		return 1
	}

	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16

	if 32 < bitCap {
		n |= n >> 32
	}

	return n + 1
}
