# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
ENV CGO_ENABLED=0 

WORKDIR /app

COPY . .
RUN go mod download

RUN go test ./internal/... -coverprofile=coverage.out -covermode=atomic

WORKDIR /app/cmd/article
RUN go build -o /article

EXPOSE 8080

CMD [ "/article" ]