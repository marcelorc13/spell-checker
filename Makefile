build: ## Build spell-checker and gen-dict binaries into bin/
	go build -o bin/spell-checker ./src/cmd/spell-checker/
	go build -o bin/gen-dict ./src/cmd/gen-dict/

run: ## Run spell-checker — usage: make run input=data/input_basico.json
	bin/spell-checker -input $(input)

gen: ## Generate stress-test dictionary (200k words, 1k queries) → data/input_estresse.json
	bin/gen-dict

test: ## Run all unit tests
	go test ./...

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"}; {printf "  %-8s %s\n", $$1, $$2}'

.PHONY: build run gen test help
