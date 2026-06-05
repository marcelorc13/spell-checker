package fuzzy

import (
	"testing"

	"github.com/marceloramalho/spell-checker/src/trie"
)

func buildTrie(words []struct{ w string; f int }) *trie.Node {
	tr := trie.New()
	for _, e := range words {
		tr.Insert(e.w, e.f)
	}
	return tr.Root()
}

func TestFuzzyExactMatch(t *testing.T) {
	root := buildTrie([]struct{ w string; f int }{{"hello", 1500}})
	results := FuzzySearch(root, "hello", 2)
	if len(results) != 1 || results[0].Word != "hello" || results[0].Dist != 0 {
		t.Fatalf("exact match failed: %+v", results)
	}
}

func TestFuzzyOneDeletion(t *testing.T) {
	root := buildTrie([]struct{ w string; f int }{{"hello", 1500}, {"help", 1200}})
	results := FuzzySearch(root, "helo", 2)
	found := map[string]bool{}
	for _, r := range results {
		found[r.Word] = true
	}
	if !found["hello"] {
		t.Fatal("expected hello in results for query helo")
	}
}

func TestFuzzyNoMatch(t *testing.T) {
	root := buildTrie([]struct{ w string; f int }{{"hello", 1500}})
	results := FuzzySearch(root, "xyz", 2)
	if len(results) != 0 {
		t.Fatalf("expected no results for xyz, got %+v", results)
	}
}

func TestFuzzyMaxDistEnforced(t *testing.T) {
	root := buildTrie([]struct{ w string; f int }{{"abcdef", 100}})
	// "abc" vs "abcdef" = 3 deletions → outside maxDist=2
	results := FuzzySearch(root, "abc", 2)
	if len(results) != 0 {
		t.Fatalf("expected no results, got %+v", results)
	}
}
