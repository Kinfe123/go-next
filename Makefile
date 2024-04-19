build:
	go build -o bin/account-go
run: build:
	./bin/account-go
test:
	go test -v ./...
localrun:
	go run main.go
