package recursion

import (
	"testing"
	"math/rand"
)

const (
	FIXED_LENGTH = 1000
	FIXED_ITERATIONS = 10000
)

func makeRandomNumberSlice(length int) []int {
	s := make([]int, 0, length)
	for range length {
		s = append(s, rand.Intn(length))
	}
	return s
}

func BenchmarkSumListWithoutRecursion(b *testing.B) {
	b.StopTimer()
	s := makeRandomNumberSlice(FIXED_LENGTH)
	b.StartTimer()
	for range b.N {
		SumListWithoutRecursion(s)
	}
}

func TestSumListWithRecursion(t *testing.T) {
	// given
	s := []int{2, 4, 6}
	expected := 12
	// when
	obtained := SumListWithRecursion(s)
	// then
	if expected != obtained {
		t.Errorf("got: %d, expected: %d", obtained, expected)
	}
}

func BenchmarkSumList(b *testing.B) {
	b.StopTimer()
	s := makeRandomNumberSlice(FIXED_LENGTH)
	b.StartTimer()
	for range b.N {
		SumListWithRecursion(s)
	}
}

func TestCountListItems(t *testing.T) {
	// given
	s := []int{2, 4, 6}
	expected := 3
	// when
	obtained := CountListItems(s)
	// then
	if expected != obtained {
		t.Errorf("got: %d, expected: %d", obtained, expected)
	}
}

func BenchmarkCountListItems(b *testing.B) {
	b.StopTimer()
	s := makeRandomNumberSlice(FIXED_LENGTH)
	b.StartTimer()
	for range b.N {
		CountListItems(s)
	}
}

func TestMax(t *testing.T) {
	// given
	s := []int{2, 4, 6, 1}
	expected := 6
	// given
	obtained := Max(s)
	// then
	if expected != obtained {
		t.Errorf("got: %d, expected: %d", obtained, expected)
	}
}
