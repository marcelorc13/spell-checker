build:
	go build -o bin/spell-checker ./src/

run:
	bin/spell-checker -input $(input)

test:
	go test ./...

.PHONY: build run test
