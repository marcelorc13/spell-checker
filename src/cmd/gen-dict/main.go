package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type DictEntry struct {
	Word string `json:"word"`
	Freq int    `json:"freq"`
}

type Input struct {
	Dictionary []DictEntry `json:"dictionary"`
	Queries    []string    `json:"queries"`
}

var syllables = []string{
	"ba", "be", "bi", "bo", "bu",
	"ca", "ce", "ci", "co", "cu",
	"da", "de", "di", "do", "du",
	"fa", "fe", "fi", "fo", "fu",
	"ga", "ge", "gi", "go", "gu",
	"la", "le", "li", "lo", "lu",
	"ma", "me", "mi", "mo", "mu",
	"na", "ne", "ni", "no", "nu",
	"pa", "pe", "pi", "po", "pu",
	"ra", "re", "ri", "ro", "ru",
	"sa", "se", "si", "so", "su",
	"ta", "te", "ti", "to", "tu",
	"va", "ve", "vi", "vo", "vu",
	"xa", "xe", "xi", "xo", "xu",
	"za", "ze", "zi", "zo", "zu",
}

// word builds a unique pronounceable word from an index by encoding it
// as a sequence of syllables (base-len(syllables)) plus a length-varying tail.
func word(idx int) string {
	n := len(syllables)
	w := ""
	// at least 2 syllables, more for higher indices to keep words varied in length
	rest := idx
	for {
		w = syllables[rest%n] + w
		rest /= n
		if rest == 0 {
			break
		}
	}
	if len(w) < 4 {
		w = syllables[idx%n] + w
	}
	return w
}

const dictSize = 200000
const querySize = 1000

func main() {
	rng := rand.New(rand.NewSource(42))

	// freqs are a shuffled permutation of 1..dictSize so every word has a
	// unique frequency -- this keeps MergeSort output deterministic even
	// though Go map iteration order (trie children, DFS) is randomized.
	freqs := make([]int, dictSize)
	for i := range freqs {
		freqs[i] = dictSize - i
	}
	rng.Shuffle(len(freqs), func(i, j int) { freqs[i], freqs[j] = freqs[j], freqs[i] })

	dict := make([]DictEntry, 0, dictSize)
	seen := make(map[string]bool, dictSize)
	words := make([]string, 0, dictSize)

	for i := 0; len(dict) < dictSize; i++ {
		w := word(i)
		if seen[w] {
			continue
		}
		seen[w] = true
		dict = append(dict, DictEntry{Word: w, Freq: freqs[len(dict)]})
		words = append(words, w)
	}

	queries := make([]string, 0, querySize)
	for q := 0; q < querySize; q++ {
		src := []rune(words[rng.Intn(len(words))])
		edits := 1 + rng.Intn(2) // 1 or 2 edits -> distance <= 2
		for e := 0; e < edits; e++ {
			if len(src) == 0 {
				break
			}
			pos := rng.Intn(len(src))
			switch rng.Intn(3) {
			case 0: // substitution
				src[pos] = randLetter(rng)
			case 1: // insertion
				r := randLetter(rng)
				src = append(src[:pos], append([]rune{r}, src[pos:]...)...)
			case 2: // deletion
				src = append(src[:pos], src[pos+1:]...)
			}
		}
		queries = append(queries, string(src))
	}

	input := Input{Dictionary: dict, Queries: queries}
	out, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "marshal:", err)
		os.Exit(1)
	}
	if err := os.WriteFile("data/input_estresse.json", out, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "write:", err)
		os.Exit(1)
	}
	fmt.Printf("generated %d dictionary terms, %d queries\n", len(dict), len(queries))
}

func randLetter(rng *rand.Rand) rune {
	return rune('a' + rng.Intn(26))
}
