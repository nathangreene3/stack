package stack

import "fmt"

// A Stack stack is a last-in-first-out (lifo) data structure.
type Stack struct {
	values []interface{}
	size   int
}

// New returns a new stack.
func New(values ...interface{}) *Stack {
	stk := Stack{
		values: append(make([]interface{}, 0, nextPow2(len(values))), values...),
		size:   len(values),
	}

	return &stk
}

// Clean frees up space that is no longer needed.
func (s *Stack) Clean() *Stack {
	if s.size < cap(s.values)>>1 {
		// 2*size < cap --> reduce cap to 2^n >= len for minimal n
		s.values = append(make([]interface{}, 0, nextPow2(s.size)), s.values[:s.size]...)
	}

	return s
}

// Clear removes all values from the stack setting the size to zero.
func (s *Stack) Clear() *Stack {
	s.size = 0
	return s
}

// Contains determines if a value is in the stack.
func (s *Stack) Contains(value interface{}) bool {
	for i := s.size - 1; 0 <= i; i-- {
		if s.values[i] == value {
			return true
		}
	}

	return false
}

// Copy returns a copy of a stack.
func (s *Stack) Copy() *Stack {
	cpy := Stack{
		values: append(make([]interface{}, 0, nextPow2(s.size)), s.values[:s.size]...),
		size:   s.size,
	}

	return &cpy
}

// Equal determines if two stacks are equal.
func (s *Stack) Equal(stack *Stack) bool {
	switch {
	case s == stack:
		return true
	case s.size != stack.size:
		return false
	}

	for i := 0; i < s.size; i++ {
		if s.values[i] != stack.values[i] {
			return false
		}
	}

	return true
}

// Peek returns the top of the stack without modifying the stack's
// state.
func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}

	return s.values[s.size-1]
}

// Pop removes and returns the top value of the stack. If the stack is
// empty, nil will be returned.
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}

	s.size--
	return s.values[s.size]
}

// Push adds a value onto the stack.
func (s *Stack) Push(values ...interface{}) *Stack {
	n := len(s.values) - s.size // Number of values that can be copied to the stack's values
	if 0 < n {
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
func (s *Stack) Size() int {
	return s.size
}

// Values returns a copy of the stack's values. The bottom of the stack
// is the zero index.
func (s *Stack) Values() []interface{} {
	return append(make([]interface{}, 0, s.size), s.values[:s.size]...)
}

// String returns a representation of a stack.
func (s *Stack) String() string {
	return fmt.Sprintf("{values: %v size: %d}", s.values[:s.size], s.size)
}

// --------------------------------------------------------------------
// Helpers
// --------------------------------------------------------------------

// nextPow2 returns the next power of two greater than or equal to a
// given number.
// 	n < 0 --> 0 *Undefined, but safe to call
// 	n = 0 --> 1
// 	n > 0 --> 2^m such that 2^m >= n for  minimal m >= 0
func nextPow2(n int) int {
	// bitCap is the number of bits in an int
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
