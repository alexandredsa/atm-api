run:
	@go run .

test:
	@go test -v -coverpkg=./... -coverprofile cover.out ./...

coverage:
	@go tool cover -html=cover.out -o cover.html	