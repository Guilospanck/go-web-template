#---- Building Stage
FROM golang:alpine as build

WORKDIR /build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build --ldflags "-s -w" -o exec main.go

#---- Image Stage
FROM scratch

WORKDIR /app
ENV GO_ENV=production
COPY --from=build ./build/exec ./
COPY ./.env.* ./
EXPOSE 4444

CMD ["./exec"]