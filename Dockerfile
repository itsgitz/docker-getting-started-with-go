# syntax=docker/dockerfile:1
FROM golang:1.19.4 AS builder
WORKDIR $GOPATH/src/github.com/itsgitz/docker-getting-started-with-go
COPY . .
RUN go mod download
RUN go install

FROM ubuntu:22.04
COPY --from=builder /go/bin/music-playlist /app/music-playlist
WORKDIR /app

ENTRYPOINT ["./music-playlist"]
EXPOSE 3000
