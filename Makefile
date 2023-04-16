NAME=mytest


.PHONY: build
build: 
	go build -mod=vendor -o $(NAME) ./cmd/example

run:
	go run ./cmd/example/main.go

test_cover:
	go test -coverprofile coverage.out -json ./... > reports.json

.PHONY: clean
clean:
	rm -rf $(NAME)
