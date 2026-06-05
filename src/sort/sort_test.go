package sort

import (
	"testing"

	"github.com/marceloramalho/spell-checker/src/fuzzy"
)

func TestMergeSortFreqDesc(t *testing.T) {
	input := []fuzzy.Suggestion{
		{Word: "cat", Freq: 100},
		{Word: "hello", Freq: 1500},
		{Word: "world", Freq: 800},
		{Word: "help", Freq: 1200},
	}
	sorted := MergeSort(input)
	expected := []int{1500, 1200, 800, 100}
	for i, s := range sorted {
		if s.Freq != expected[i] {
			t.Fatalf("pos %d: got freq %d, want %d", i, s.Freq, expected[i])
		}
	}
}

func TestMergeSortEmpty(t *testing.T) {
	result := MergeSort(nil)
	if len(result) != 0 {
		t.Fatalf("expected empty, got %+v", result)
	}
}

func TestMergeSortSingle(t *testing.T) {
	input := []fuzzy.Suggestion{{Word: "x", Freq: 42}}
	result := MergeSort(input)
	if len(result) != 1 || result[0].Freq != 42 {
		t.Fatalf("single item sort failed: %+v", result)
	}
}
