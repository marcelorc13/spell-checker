package sort

import "github.com/marceloramalho/spell-checker/src/fuzzy"

func MergeSort(items []fuzzy.Suggestion) []fuzzy.Suggestion {
	if len(items) <= 1 {
		return items
	}
	mid := len(items) / 2
	left := MergeSort(items[:mid])
	right := MergeSort(items[mid:])
	return merge(left, right)
}

func merge(left, right []fuzzy.Suggestion) []fuzzy.Suggestion {
	result := make([]fuzzy.Suggestion, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Freq >= right[j].Freq {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
