run:
	go build -o bin/main cmd/api/main.go && ./bin/main


test:
	go test ./... | { grep -v 'no test files'; true; }
