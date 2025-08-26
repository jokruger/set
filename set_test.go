package set

import (
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	s := NewFromElements(1, 2, 2, 3)
	if s.Len() != 3 {
		t.Error("Set should have length 3")
	}

	s = NewFromSlice([]int{1, 2, 2, 3})
	if s.Len() != 3 {
		t.Error("Set should have length 3")
	}

	sl1 := []int{1, 2, 2, 3}
	s = NewFromSeq(slices.Values(sl1))
	if s.Len() != 3 {
		t.Error("Set should have length 3")
	}

	s = NewFromSet(NewFromElements(1, 2, 2, 3))
	if s.Len() != 3 {
		t.Error("Set should have length 3")
	}
}

func TestInts(t *testing.T) {
	s := New[int](WithCapacity(10))
	if s.Contains(42) {
		t.Error("Set should not contain 42")
	}
	s.Add(42)
	if !s.Contains(42) {
		t.Error("Set should contain 42")
	}
	s.Remove(42)
	if s.Contains(42) {
		t.Error("Set should not contain 42")
	}
}

func TestFromIntsSlice(t *testing.T) {
	s := NewFromSlice([]int{1, 2, 3})
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Error("Set should contain 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}
	if !s.Contains(2, 3) {
		t.Error("Set should contain 2 and 3")
	}
	if s.Contains(2, 4) {
		t.Error("Set should not contain 4")
	}
	if !s.ContainsAny(2, 4) {
		t.Error("Set should contain 2")
	}
	if s.ContainsAny(4, 5) {
		t.Error("Set should not contain 4 or 5")
	}

	s = New[int]()
	s.AddMany(1, 2, 3)
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Error("Set should contain 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}
}

func TestFromIntsSet(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(2, 3, 4)
	s := NewFromSet(s1)
	s.AddSet(s2)
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Error("Set should contain 3")
	}
	if !s.Contains(4) {
		t.Error("Set should contain 4")
	}
	if s.Contains(5) {
		t.Error("Set should not contain 5")
	}
}

func TestFromIntsIterator(t *testing.T) {
	slice := []int{1, 2, 3}
	s := NewFromSeq(slices.Values(slice))
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Error("Set should contain 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}
}

func TestAddInts(t *testing.T) {
	s := New[int]()
	s.AddMany(1, 2, 3)
	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
	if !s.Contains(2) {
		t.Error("Set should contain 2")
	}
	if !s.Contains(3) {
		t.Error("Set should contain 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}

	s.AddSet(NewFromElements(4, 5, 6))
	if !s.Contains(4) {
		t.Error("Set should contain 4")
	}
	if !s.Contains(5) {
		t.Error("Set should contain 5")
	}
	if !s.Contains(6) {
		t.Error("Set should contain 6")
	}

	s.AddSlice([]int{7, 8, 9})
	if !s.Contains(7) {
		t.Error("Set should contain 7")
	}
	if !s.Contains(8) {
		t.Error("Set should contain 8")
	}
	if !s.Contains(9) {
		t.Error("Set should contain 9")
	}

	slice := []int{10, 11, 12}
	s.AddSeq(slices.Values(slice))
	if !s.Contains(10) {
		t.Error("Set should contain 10")
	}
	if !s.Contains(11) {
		t.Error("Set should contain 11")
	}
	if !s.Contains(12) {
		t.Error("Set should contain 12")
	}
}

func TestRemoveInts(t *testing.T) {
	s := NewFromElements(1, 2, 3, 4, 5, 6, 7, 9)
	s.RemoveMany(1, 2)
	if s.Contains(1) {
		t.Error("Set should not contain 1")
	}
	if s.Contains(2) {
		t.Error("Set should not contain 2")
	}

	s.RemoveSet(NewFromElements(3, 4))
	if s.Contains(3) {
		t.Error("Set should not contain 3")
	}
	if s.Contains(4) {
		t.Error("Set should not contain 4")
	}

	s.RemoveSlice([]int{5, 6})
	if s.Contains(5) {
		t.Error("Set should not contain 5")
	}
	if s.Contains(6) {
		t.Error("Set should not contain 6")
	}

	slice := []int{7, 9}
	s.RemoveSeq(slices.Values(slice))
	if s.Contains(7) {
		t.Error("Set should not contain 7")
	}
	if s.Contains(9) {
		t.Error("Set should not contain 9")
	}
}

func TestClearInts(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	if s.IsEmpty() {
		t.Error("Set should not be empty")
	}
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Set should be empty")
	}
}

func TestContainsInts(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	if !s.Contains(1, 2) {
		t.Error("Set should contain 1 and 2")
	}
	if s.Contains(1, 4) {
		t.Error("Set should not contain 4")
	}
	if !s.ContainsSet(NewFromElements(1, 2)) {
		t.Error("Set should contain 1 and 2")
	}
	if s.ContainsSet(NewFromElements(1, 4)) {
		t.Error("Set should not contain 4")
	}
	if !s.ContainsSlice([]int{1, 2}) {
		t.Error("Set should contain 1 and 2")
	}
	if s.ContainsSlice([]int{1, 4}) {
		t.Error("Set should not contain 4")
	}

	slice := []int{1, 2}
	if !s.ContainsSeq(slices.Values(slice)) {
		t.Error("Set should contain 1 and 2")
	}
	slice = []int{1, 4}
	if s.ContainsSeq(slices.Values(slice)) {
		t.Error("Set should not contain 4")
	}
}

func TestContainsAnyInts(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	if !s.ContainsAny(1, 4) {
		t.Error("Set should contain 1")
	}
	if s.ContainsAny(4, 5) {
		t.Error("Set should not contain 4 or 5")
	}
	if !s.ContainsAnyFromSet(NewFromElements(1, 4)) {
		t.Error("Set should contain 1")
	}
	if s.ContainsAnyFromSet(NewFromElements(4, 5)) {
		t.Error("Set should not contain 4 or 5")
	}
	if !s.ContainsAnyFromSlice([]int{1, 4}) {
		t.Error("Set should contain 1")
	}
	if s.ContainsAnyFromSlice([]int{4, 5}) {
		t.Error("Set should not contain 4 or 5")
	}

	slice := []int{1, 4}
	if !s.ContainsAnyFromSeq(slices.Values(slice)) {
		t.Error("Set should contain 1")
	}
	slice = []int{4, 5}
	if s.ContainsAnyFromSeq(slices.Values(slice)) {
		t.Error("Set should not contain 4 or 5")
	}
}

func TestLen(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	if s.Len() != 3 {
		t.Error("Set should have length 3")
	}
}

func TestEqualInts(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3)
	if !s1.Equal(s2) {
		t.Error("Sets should be equal")
	}
	s3 := NewFromElements(1, 2, 3, 4)
	if s1.Equal(s3) {
		t.Error("Sets should not be equal")
	}
}

func TestCloneInts(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := s1.Clone()
	if !s1.Equal(s2) {
		t.Error("Sets should be equal")
	}
}

func TestToSlice(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Error("Slice should have length 3")
	}
	slices.Sort(slice)
	if slice[0] != 1 {
		t.Error("Slice should contain 1")
	}
	if slice[1] != 2 {
		t.Error("Slice should contain 2")
	}
	if slice[2] != 3 {
		t.Error("Slice should contain 3")
	}
}

func TestToSeq(t *testing.T) {
	s := NewFromElements(1, 2, 3)
	for e := range s.ToSeq() {
		if !s.Contains(e) {
			t.Errorf("Set should contain %d", e)
		}
	}
}

func TestIsSubset(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3, 4)
	if !s1.IsSubsetOf(s2) {
		t.Error("Set should be a subset of other set")
	}
	if s2.IsSubsetOf(s1) {
		t.Error("Set should not be a subset of other set")
	}
}

func TestIsSuperset(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3, 4)
	if s1.IsSupersetOf(s2) {
		t.Error("Set should not be a superset of other set")
	}
	if !s2.IsSupersetOf(s1) {
		t.Error("Set should be a superset of other set")
	}
}

func TestIsProperSubset(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3, 4)
	if !s1.IsProperSubsetOf(s2) {
		t.Error("Set should be a proper subset of other set")
	}
	if s2.IsProperSubsetOf(s1) {
		t.Error("Set should not be a proper subset of other set")
	}
	s3 := NewFromElements(1, 2, 3)
	if s1.IsProperSubsetOf(s3) {
		t.Error("Set should not be a proper subset of other set")
	}
}

func TestIsProperSuperset(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3, 4)
	if s1.IsProperSupersetOf(s2) {
		t.Error("Set should not be a proper superset of other set")
	}
	if !s2.IsProperSupersetOf(s1) {
		t.Error("Set should be a proper superset of other set")
	}
	s3 := NewFromElements(1, 2, 3)
	if s1.IsProperSupersetOf(s3) {
		t.Error("Set should not be a proper superset of other set")
	}
}

func TestDiff(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(1, 2, 3, 4)
	diff1 := s1.Diff(s2)
	if diff1.Len() != 0 {
		t.Error("Diff should be empty")
	}
	diff2 := s2.Diff(s1)
	if diff2.Len() != 1 {
		t.Error("Diff should have length 1")
	}
	if !diff2.Contains(4) {
		t.Error("Diff should contain 4")
	}
}

func TestSymmetricDiff(t *testing.T) {
	s1 := NewFromElements(1, 2, 3, 4, 5)
	s2 := NewFromElements(2, 4, 6)
	diff := s1.SymmetricDiff(s2)
	if diff.Len() != 4 {
		t.Error("Diff should have length 4")
	}
	if !diff.Contains(1) {
		t.Error("Diff should contain 1")
	}
	if !diff.Contains(3) {
		t.Error("Diff should contain 3")
	}
	if !diff.Contains(5) {
		t.Error("Diff should contain 5")
	}
	if !diff.Contains(6) {
		t.Error("Diff should contain 6")
	}
}

func TestUnion(t *testing.T) {
	s1 := NewFromElements(1, 2, 3)
	s2 := NewFromElements(2, 3, 4)
	union := s1.Union(s2)
	if union.Len() != 4 {
		t.Error("Union should have length 4")
	}
	if !union.Contains(1) {
		t.Error("Union should contain 1")
	}
	if !union.Contains(2) {
		t.Error("Union should contain 2")
	}
	if !union.Contains(3) {
		t.Error("Union should contain 3")
	}
	if !union.Contains(4) {
		t.Error("Union should contain 4")
	}
}

func TestFilter(t *testing.T) {
	s := NewFromElements(1, 2, 3, 4, 5)
	filtered := s.Filter(func(e int) bool {
		return e%2 == 0
	})
	if filtered.Len() != 2 {
		t.Error("Filtered set should have length 2")
	}
	if !filtered.Contains(2) {
		t.Error("Filtered set should contain 2")
	}
	if !filtered.Contains(4) {
		t.Error("Filtered set should contain 4")
	}
}
