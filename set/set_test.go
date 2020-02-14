package set

import "testing"

func TestAdd(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)

	expected := []interface{}{1, 2}
	got := s.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
func TestDelete(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Delete(1)

	expected := []interface{}{1, 2}
	got := s.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestContains(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)

	testCases := []struct {
		desc     string
		v        int
		expected bool
		got      bool
	}{
		{
			desc:     "contain",
			v:        1,
			expected: true,
		},
		{
			desc:     "does not contain",
			v:        0,
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := s.Contains(tC.v)
			if tC.expected != got {
				t.Errorf("expected %v, got %v", tC.expected, got)
			}
		})
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
