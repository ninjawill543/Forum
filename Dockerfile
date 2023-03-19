FROM golang:1.19.7-alpine3.17

WORKDIR /src

RUN apk add build-base
RUN apk add --no-cache --upgrade bash

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN chmod +x start.sh
EXPOSE 8080

ENTRYPOINT ["/src/start.sh"]