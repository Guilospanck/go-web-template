# Go Web Template
This is a template to be used when writing Go for the web.

It covers concepts of Clean Code, unit testing, GitHub actions and other practices and, for those, it uses:
- [Gin](https://github.com/gin-gonic/gin) to handle HTTP requests;
- [Zap](https://github.com/uber-go/zap) for logging;
- [GoDotEnv](https://github.com/joho/godotenv) for managing environments; and
- [Testify](https://github.com/stretchr/testify) for unit testings.

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

## Understanding the format of the application
Every package tries to follow the pattern of Clean Code architecture.

The entry point of the program is the `main.go` file as usual. It'll execute the commands that are inside `cmd/root.go` and panic if something goes wrong.

The `cmd/root.go` file is responsible for instantiating all the packages necessaries for running the application (the "container" that resides inside the `cmd/iot.go` file) and start the HTTP server.

As everything inside the `pkg` folder follows the concept of Clean Code, the packages only communicate with each other via `interfaces`. Therefore `cmd/iot.go` is where the concrete implementations of those interfaces are instantiated.

Finally, the `pkg` folder contains the following structure:
```
.
├── application
│   ├── errors
│   ├── interfaces
│   └── usecases
│       └── ping
├── domain
│   ├── dtos
│   └── usecases
├── infrastructure
│   ├── adapters
│   ├── environments
│   ├── http_server
│   └── logger
└── interface
    └── http
        ├── factories
        ├── handlers
        └── presenters
```
And the flux happens like this:
```
interface <-> applications <-> infrastructure/domain
```
- `interface`: defines the ways that our application can interface with the outside world. In this case, there's only `http`, but there may exist `graphql`, `gRPC` or any other type of communication.
    - `presenters`: here we define our routes (endpoints)
    - `handlers`: they act like controllers, changing the traffic from outside the application to the right path inside it.
- `application`: here lies our "use cases". They describe what the application does in a higher level. The point is that the use cases should explain to an outsider, just by looking at their names, what is the flux and functions that the application have.
- `infrastructure`: here usually lies the libraries/frameworks/ORMs that we inject in our application. This is the "external layer" of the architecture. They should allow easy replacement without compromising the whole application. We should not depend on them.
- `domain`: everything related to our business cases. DTOS, Enums, the interfaces that represent our use cases.

## Running tests
To run tests, you can run the commands:
```bash
make test
# or
make test-cov
```