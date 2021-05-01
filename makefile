run:
	go build -o bin/main cmd/api/main.go && DB_PASS=test ./bin/main
cover: 
	go test -coverpkg ./... ./... -coverprofile=coverage.out

cover-html:
	go tool cover -html=coverage.out

test:
	go test github.com/brianfromlife/golang-ecs/tests --coverpkg ./... ./...

build:
	go build -o main cmd/api/main.go