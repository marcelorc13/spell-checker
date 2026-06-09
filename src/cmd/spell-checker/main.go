package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	appfuzzy "github.com/marceloramalho/spell-checker/src/fuzzy"
	appio "github.com/marceloramalho/spell-checker/src/io"
	appsort "github.com/marceloramalho/spell-checker/src/sort"
	apptrie "github.com/marceloramalho/spell-checker/src/trie"
)

func main() {
	defaultInputPath := "data/input_basico.json"
	inputPath := flag.String("input", defaultInputPath, "input file")
	flag.Parse()

	fmt.Printf("\ninput: %s\n", *inputPath)

	t0 := time.Now()
	input, err := appio.ReadInput(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("loaded %d dictionary entries, %d queries \n",
		len(input.Dictionary), len(input.Queries))

	t := apptrie.New()
	for _, entry := range input.Dictionary {
		t.Insert(entry.Word, entry.Freq)
	}

	output := &appio.Output{}
	for _, query := range input.Queries {
		suggestions := appfuzzy.FuzzySearch(t.Root(), query, 2)
		suggestions = appsort.MergeSort(suggestions)

		words := make([]string, len(suggestions))
		for j, s := range suggestions {
			words[j] = s.Word
		}
		output.Results = append(output.Results, appio.Result{
			Query:       query,
			Suggestions: words,
		})
	}
	fmt.Printf("fuzzy search done \n")

	dp := "out"
	fn := strings.SplitAfter(*inputPath, "_")[1]

	outputPath := fmt.Sprintf("%s/output_%s", dp, fn)

	if _, err := os.Stat(dp); os.IsNotExist(err) {
		err := os.Mkdir(dp, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	fmt.Printf("writing output: %s\n", outputPath)
	if err := appio.WriteOutput(outputPath, output); err != nil {
		fmt.Fprintf(os.Stderr, "write output: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("done. \ntotal: %s\n", time.Since(t0))
}
