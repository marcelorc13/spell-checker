package io

import (
	"encoding/json"
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

type Result struct {
	Query       string   `json:"query"`
	Suggestions []string `json:"suggestions"`
}

type Output struct {
	Results []Result `json:"results"`
}

func ReadInput(path string) (*Input, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var input Input
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	return &input, nil
}

func WriteOutput(path string, output *Output) error {
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
