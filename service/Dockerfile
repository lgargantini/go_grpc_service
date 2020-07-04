FROM golang:1.13-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd server && go build server.go

CMD [ "./server/server" ]
