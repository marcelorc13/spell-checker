build:
	go build -o bin/spell-checker ./src/

run:
	bin/spell-checker -input data/input_basico.json -output data/output.json

test:
	go test ./...

.PHONY: build run test
