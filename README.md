# Go Web Template
This is a template to be used when writing Go for the web.

It covers concepts of Clean Code, unit testing, GitHub actions and other practices.

## Installation
Be sure to have [Go](https://go.dev/doc/install) installed. 

Then, clone this repository:
```bash
git clone https://github.com/Guilospanck/go-web-template.git
cd go-web-template/
```
And then install dependencies:
```bash
go mod download
```

## How to use
If you wanna dive into the "project" and play with it, you can just run:
```bash
make run-local
```

But, on the other hand, if you wanna use it as a template for your next Go project, just click in `Use this template`.

## Playing with the application
### Endpoints
There's only one endpoint for this sample application. It will listen at `http://localhost:4444/ping` for a POST request and will return the following JSON:
```json
{
	"Pong": 1
}
```

### Curl
You can also make the requests using `curl`:
```bash
curl --request POST \
  --url http://localhost:4444/ping \
  --header 'Content-Type: application/json' \
  --data '{}'
```

## Running tests
To run tests, you can run the commands:
```bash
make test
# or
make test-cov
```