FROM golang:1.17.10-alpine AS build

ARG VERSION=0.0.1
ARG BIN_NAME=smtp-sender

WORKDIR /app
COPY ./src .
COPY .env .env

RUN go mod vendor
RUN CGO_ENABLED=0 go build -ldflags="-X github.com/OrlangurDux/rest-api-boilerplate-with-smtp.Version=${VERSION}" -o "${BIN_NAME}" ./

FROM scratch AS bin
ARG BIN_NAME=smtp-sender

COPY --from=build /app/${BIN_NAME} /${BIN_NAME}
COPY --from=build /app/.env /.env

ENTRYPOINT [ "/smtp-sender" ]

EXPOSE 8025
