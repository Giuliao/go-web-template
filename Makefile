
build:
	go build -o bin/app

run: build
	./bin/app

test:
	go test -v 

swag: 
	swag init -g ./cmd/main.go -o ./api