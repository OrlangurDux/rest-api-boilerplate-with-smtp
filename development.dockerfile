FROM golang:latest

ENV GO111MODULE=on
ENV PORT=8025
WORKDIR /app
COPY ./src/go.mod /app
COPY ./src/go.sum /app

RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon@latest
COPY ./src /app
ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main
