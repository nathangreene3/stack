package stack

import "testing"

func TestNextPow2(t *testing.T) {
	expNextPow2 := map[int]int{
		-1: 0, // Undefined case: n < 0

		0: 1,
		1: 1,

		2: 2,

		3: 4,
		4: 4,

		5: 8,
		7: 8,
		8: 8,

		9:  16,
		15: 16,
		16: 16,

		17: 32,
		31: 32,
		32: 32,

		33: 64,
		63: 64,
		64: 64,

		65:  128,
		127: 128,
		128: 128,

		129: 256,
		255: 256,
		256: 256,
	}

	for n, exp := range expNextPow2 {
		if rec := nextPow2(n); exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", exp, rec)
		}
	}
}

func TestStack(t *testing.T) {
	var (
		values  = []interface{}{0, 1, 2, 3, 4}
		expSize = len(values) >> 1
		stk     = New(values[:expSize]...)
	)

	if rec := stk.Size(); rec != expSize {
		t.Errorf("\nexpected %d\nreceived %d\n", expSize, rec)
		return
	}

	for _, value := range values[len(values)>>1:] {
		stk.Push(value)
		expSize++
		if rec := stk.Size(); expSize != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", expSize, rec)
			return
		}

		if rec := stk.Peek(); value != rec {
			t.Errorf("\nexpected %v\nreceived %v\n", values[expSize], value)
		}
	}

	for _, value := range values {
		if !stk.Contains(value) {
			t.Errorf("\nexpected %v to be contained\n", value)
		}
	}

	for i, value := range stk.Values() {
		if values[i] != value {
			t.Errorf("\nexpected %v\nreceived %v\n", values[i], value)
		}
	}

	cpy := stk.Copy()
	if !stk.Equal(cpy) {
		t.Errorf("\nexpected %v\nreceived %v\n", stk, cpy)
	}

	for expSize != 0 {
		value := stk.Pop()
		expSize--
		if value != values[expSize] {
			t.Errorf("\nexpected %v\nreceived %v\n", values[expSize], value)
			return
		}

		if rec := stk.Size(); expSize != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", expSize, rec)
			return
		}

		if stk.Contains(value) {
			t.Errorf("\nexpected %v to not be contained\n", value)
		}
	}

	stk.Clean()
	if len(stk.values) != 0 || cap(stk.values) != 1 {
		t.Errorf("\nexpected len %d, cap %d\nreceived %d, %d\n", 0, 1, len(stk.values), cap(stk.values))
	}

	if rec := stk.Peek(); rec != nil {
		t.Errorf("\nexpected %v\nreceived %v\n", nil, rec)
	}

	if rec := stk.Pop(); rec != nil {
		t.Errorf("\nexpected %v\nreceived %v\n", nil, rec)
	}
}
