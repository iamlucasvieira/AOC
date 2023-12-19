package utils

import "testing"

func TestStack_Push(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s) != 3 {
		t.Errorf("Push: expected len(s) to be 3, got %d", len(s))
	}

	if s[0] != 1 {
		t.Errorf("Push: expected s[0] to be 1, got %d", s[0])
	}

	if s[1] != 2 {
		t.Errorf("Push: expected s[1] to be 2, got %d", s[1])
	}

	if s[2] != 3 {
		t.Errorf("Push: expected s[2] to be 3, got %d", s[2])
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	testCases := []struct {
		id     string
		expect int
	}{
		{
			id:     "first",
			expect: 3,
		},
		{
			id:     "second",
			expect: 2,
		},
		{
			id:     "third",
			expect: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.id, func(t *testing.T) {
			v, ok := s.Pop()
			if !ok {
				t.Errorf("Pop: expected ok to be true, got false")
			}

			if v != tc.expect {
				t.Errorf("Pop: expected v to be %d, got %d", tc.expect, v)
			}
		})
	}
}

func TestStack_Pop_Empty(t *testing.T) {
	s := Stack[int]{}

	v, ok := s.Pop()
	if ok {
		t.Errorf("Pop: expected ok to be false, got true")
	}

	if v != 0 {
		t.Errorf("Pop: expected v to be 0, got %d", v)
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, ok := s.Peek()
	if !ok {
		t.Errorf("Peek: expected ok to be true, got false")
	}

	if v != 3 {
		t.Errorf("Peek: expected v to be 3, got %d", v)
	}

	if len(s) != 3 {
		t.Errorf("Peek: expected len(s) to be 3, got %d", len(s))
	}
}
