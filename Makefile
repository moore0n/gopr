.PHONY: build

build:
	@GOARCH=amd64 GOOS=darwin go build -o ./gopr ./cmd/gopr/main.go

