build:
	@go build -o fcal .

clean:
	rm fcal

test:
	@go test ./...
