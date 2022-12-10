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

run-build:
	GO_ENV=production GOOS=linux GOARCH=amd64 go build -o main main.go   