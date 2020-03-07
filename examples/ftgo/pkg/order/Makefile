protoc=protoc --go_out=. */*.proto

.PHONY: build
build:
	go build -o ./bin/app ./cmd/app/main.go

.PHONY: run
run:
	reflex -r '(\.go|go\.mod)' -s go run ./cmd/app/main.go

.PHONY: setup
setup:
	go run ./cmd/setup/main.go
	

.PHONY: test-watch
test-watch:
	reflex -r '(\.go)' -s go test ./...
