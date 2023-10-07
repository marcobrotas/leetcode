build-parser:
	go build -o parser ./cmd/parser/bench_results_parser.go && \
	chmod +x parser

bench:
	go test -bench=. -benchmem ./... | gobenchdata --json bench.json \

parse:
	./parser

run: bench parse
