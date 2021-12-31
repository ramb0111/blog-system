# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/article
RUN go build -o /article

EXPOSE 8080

CMD [ "/article" ]