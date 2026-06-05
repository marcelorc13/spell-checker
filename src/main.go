package main

import (
	"flag"
	"fmt"
	"os"

	appfuzzy "github.com/marceloramalho/spell-checker/src/fuzzy"
	appio "github.com/marceloramalho/spell-checker/src/io"
	appsort "github.com/marceloramalho/spell-checker/src/sort"
	apptrie "github.com/marceloramalho/spell-checker/src/trie"
)

func main() {
	inputPath := flag.String("input", "data/input_basico.json", "input file")
	outputPath := flag.String("output", "data/output.json", "output file")
	flag.Parse()

	input, err := appio.ReadInput(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}

	t := apptrie.New()
	for _, entry := range input.Dictionary {
		t.Insert(entry.Word, entry.Freq)
	}

	output := &appio.Output{}
	for _, query := range input.Queries {
		suggestions := appfuzzy.FuzzySearch(t.Root(), query, 2)
		suggestions = appsort.MergeSort(suggestions)

		words := make([]string, len(suggestions))
		for i, s := range suggestions {
			words[i] = s.Word
		}
		output.Results = append(output.Results, appio.Result{
			Query:       query,
			Suggestions: words,
		})
	}

	if err := appio.WriteOutput(*outputPath, output); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}
}
