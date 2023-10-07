build-parser:
	go build -o benchmark-parser ./cmd/parser/bench_results_parser.go && \
	chmod +x benchmark-parser

bench:
	go test -bench=. -benchmem ./... | gobenchdata --json bench.json \

parse:
	./benchmark-parser

run: bench parse
