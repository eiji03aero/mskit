protoc=protoc --go_out=. */*.proto

.PHONY: build
build:
	go build -o ./bin/app ./cmd/app/main.go

.PHONY: test-watch
test-watch:
	while true; do go test -v ./...; sleep 5s; done;
