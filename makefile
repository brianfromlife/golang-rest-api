run:
	go build -o main && ./main
cover: 
	go test -coverpkg ./... ./... -coverprofile=coverage.out

cover-html:
	go tool cover -html=coverage.out

test:
	go test github.com/brianfromlife/golang-ecs/tests --coverpkg ./... ./...
