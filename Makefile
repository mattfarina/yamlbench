.PHONY: setup
setup:
	dep ensure
	go run ./cmd/gen/main.go

.PONY: bench
bench:
	go test -bench=. -benchmem
