# Spell Checker with Fuzzy Search

Academic assignment — "Projeto 6: Corretor Ortográfico com Busca Aproximada". Pure Go. No high-level algo libs.

## Required structures (implement from scratch)

- Trie: compressed prefix tree, `map[rune]*Node`
- Fuzzy DFS: Levenshtein row-carry + pruning (≤2 edits)
- MergeSort: by word freq desc — NO `sort.Slice`/`sort.Sort`/`slices.SortFunc`

## Layout

```
/
├── src/
│   ├── main.go
│   ├── trie/
│   │   ├── trie.go       # Trie struct, Insert O(L), Search O(L)
│   │   └── node.go       # map[rune]*Node, isEnd, freq, word
│   ├── fuzzy/
│   │   └── search.go     # DFS + Levenshtein row-carry + pruning
│   ├── sort/
│   │   └── sort.go       # MergeSort — hand-rolled
│   └── io/
│       └── codec.go      # JSON read/write (encoding/json OK)
├── data/
│   ├── input_basico.json
│   ├── input_avancado.json
│   ├── input_estresse.json
│   └── output_esperado_*.json
├── tools/
│   └── gen_dict.go       # converts raw word freq list → JSON inputs
├── docs/
│   └── complexity.md
├── Makefile
└── README.md
```

## IO schemas

Input:
```json
{"dictionary": [{"word": "hello", "freq": 1500}], "queries": ["helo"]}
```

Output:
```json
{"results": [{"query": "helo", "suggestions": ["hello"]}]}
```

Suggestions = words within edit distance ≤2, sorted freq desc.

## Key algorithms

### Trie node

```go
type Node struct {
    children map[rune]*Node
    isEnd    bool
    freq     int
    word     string
}
```

### Fuzzy DFS pruning

```go
// carry previous DP row through DFS
// prune when min(currentRow) > maxDist (2)
// collect when node.isEnd && currentRow[len(query)] <= maxDist
func dfs(node *trie.Node, query string, prevRow []int, maxDist int, results *[]Suggestion)
```

### MergeSort

```go
func MergeSort(items []Suggestion) []Suggestion   // O(N log N)
func merge(left, right []Suggestion) []Suggestion // sort by Freq desc
```

## Complexity (for docs/complexity.md)

| Op | Time | Space |
|----|------|-------|
| Trie Insert | O(L) | O(N·L) worst |
| FuzzySearch DFS | O(N·L) worst, O(b^d · L) avg w/ pruning | — |
| MergeSort | O(N log N) | O(N) |

## Anti-fraud constraints

- No `sort.Slice`, `sort.Sort`, `slices.SortFunc`, `container/heap`, `container/list`
- Trie must be `map[rune]*Node` — not `map[string]any`
- Levenshtein computed dynamically per query — not hardcoded

## Build

```makefile
build:
    go build -o spell-checker ./src/
run:
    ./spell-checker -input data/input_basico.json -output data/output.json
test:
    go test ./...
```

## Verify

```bash
make build
./spell-checker -input data/input_basico.json -output /tmp/out.json
diff /tmp/out.json data/output_esperado_basico.json
go test ./src/trie/ -v
go test ./src/fuzzy/ -v
go test ./src/sort/ -v
```
