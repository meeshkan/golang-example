# https://stackoverflow.com/a/55853387/10561443
FROM golang:latest

WORKDIR /app

# ENV GO111MODULE=on
# WORKDIR /go/src/github.com/unmock/example-golang/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD ["go", "run", "get.go"]
