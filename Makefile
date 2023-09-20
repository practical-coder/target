.PHONY: build clean fmt scheck vet
.DEFAULT_GOAL: build
build: test
	go build -o tg -ldflags '-s -w'
clean:
	rm -f tg
fmt:
	go fmt ./...
scheck: vet
	staticcheck ./...
test: scheck
	go test ./...
vet: fmt
	go vet ./...

