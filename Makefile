build:
	@go build -o bin/sysact cmd/sysact/main.go

test:
	@go test -v ./...

run: build
	@./bin/sysact
