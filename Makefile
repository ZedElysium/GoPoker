build:
	@go build -o GoPoker

run: build
	./GoPoker

test:
	go test -v ./...