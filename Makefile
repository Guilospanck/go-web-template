test:
	if ! [ -d "coverage" ]; then \
		echo "Creating coverage folder" ; \
		mkdir coverage; \
	fi
	go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic

test-cov:
	if ! [ -d "coverage" ]; then \
		echo "Creating coverage folder" ; \
		mkdir coverage; \
	fi
	go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic && go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

run-local:
	GO_ENV=development go run .      

build-linux:
	rm -f main  
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags "-s -w" -o main main.go

build-darwin:
	rm -f main  
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build --ldflags "-s -w" -o main main.go 