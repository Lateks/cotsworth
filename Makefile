fcal:
	@go build ./cmd/fcal

install:
	@go install ./cmd/fcal

clean:
	rm fcal

test:
	@go test ./...
