.PHONY: build

build:
	@GOARCH=amd64 GOOS=darwin go build -o ./pr ./cmd/pr/main.go

