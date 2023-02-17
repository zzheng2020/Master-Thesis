# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy
RUN go get github.com/gin-gonic/gin

COPY *.go ./

RUN go build -o /go-docker-test

EXPOSE 8080

CMD [ "/go-docker-test" ]