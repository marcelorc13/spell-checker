package fuzzy

import "github.com/marceloramalho/spell-checker/src/trie"

type Suggestion struct {
	Word string
	Freq int
	Dist int
}

func FuzzySearch(root *trie.Node, query string, maxDist int) []Suggestion {
	runes := []rune(query)
	initialRow := make([]int, len(runes)+1)
	for i := range initialRow {
		initialRow[i] = i
	}
	var results []Suggestion
	for ch, child := range root.Children {
		dfs(child, ch, runes, initialRow, maxDist, &results)
	}
	return results
}

func dfs(node *trie.Node, ch rune, query []rune, prevRow []int, maxDist int, results *[]Suggestion) {
	n := len(query)
	currRow := make([]int, n+1)
	currRow[0] = prevRow[0] + 1

	for i := 1; i <= n; i++ {
		insertCost := currRow[i-1] + 1
		deleteCost := prevRow[i] + 1
		var replaceCost int
		if query[i-1] == ch {
			replaceCost = prevRow[i-1]
		} else {
			replaceCost = prevRow[i-1] + 1
		}
		currRow[i] = min3(insertCost, deleteCost, replaceCost)
	}

	if node.IsEnd && currRow[n] <= maxDist {
		*results = append(*results, Suggestion{
			Word: node.Word,
			Freq: node.Freq,
			Dist: currRow[n],
		})
	}

	// prune: if min of currRow > maxDist, no child can improve
	minVal := currRow[0]
	for _, v := range currRow[1:] {
		if v < minVal {
			minVal = v
		}
	}
	if minVal > maxDist {
		return
	}

	for nextCh, child := range node.Children {
		dfs(child, nextCh, query, currRow, maxDist, results)
	}
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
