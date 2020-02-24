protoc=protoc --go_out=. */*.proto

build:
	go build -o ./bin/app ./cmd/app/main.go

run:
	./bin/app

protoc:
	$(protoc)

test-watch:
	while true; do go test -v ./...; sleep 5s; done;
